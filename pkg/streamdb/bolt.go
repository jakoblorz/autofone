package streamdb

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/boltdb/bolt"
)

type DebounceMode string

const (
	DebounceModeDelay      DebounceMode = "delay"
	DebounceModeActive     DebounceMode = "active"
	DebounceModeAggressive DebounceMode = "aggressive"
)

const (
	DebounceModeDelay_Interval      time.Duration = 1 * time.Minute
	DebounceModeActive_Interval     time.Duration = 4 * time.Second
	DebounceModeAggressive_Interval time.Duration = 250 * time.Millisecond
)

type Handle interface {
	Batch(func(tx *bolt.Tx) error) error
	Update(func(*bolt.Tx) error) error
	View(func(*bolt.Tx) error) error
}

type handle struct {
	*bolt.DB
	w *stream
}

func (h *handle) Batch(fn func(tx *bolt.Tx) error) error {
	defer h.w.notify()
	return h.DB.Batch(fn)
}

func (h *handle) Update(fn func(*bolt.Tx) error) error {
	defer h.w.notify()
	return h.DB.Update(fn)
}

func (h *handle) View(fn func(*bolt.Tx) error) error {
	return h.DB.View(fn)
}

type FileWriter interface {
	WriteFrom(context.Context, <-chan string)
}

type I interface {
	Get() Handle
	Put(Handle)
	Close() error
}

func Open(path string, fw FileWriter, debounceMode DebounceMode) (I, error) {
	ctx, cancel := context.WithCancel(context.Background())
	w := &stream{
		Path: path,

		mx:           new(sync.RWMutex),
		ctx:          ctx,
		cancel:       cancel,
		handleWg:     new(sync.WaitGroup),
		debounceMx:   new(sync.Mutex),
		debounceMode: debounceMode,
	}
	if err := w.open(); err != nil {
		return nil, err
	}

	if fw != nil {
		w.fileC = make(chan string, 1)
		go fw.WriteFrom(ctx, w.fileC)
	}

	return w, nil
}

type stream struct {
	mx     *sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc

	handleDb *bolt.DB
	handleWg *sync.WaitGroup

	debounceMode  DebounceMode
	debounceTimer *time.Timer
	debounceMx    *sync.Mutex

	fileC chan string

	Path string
}

func (i *stream) notify() {
	i.debounceMx.Lock()
	defer i.debounceMx.Unlock()

	if i.debounceMode == DebounceModeDelay && i.debounceTimer == nil {
		i.debounceTimer = time.AfterFunc(DebounceModeDelay_Interval, i.rotate)
		return
	}

	if i.debounceTimer != nil {
		i.debounceTimer.Stop()
	}
	if i.debounceMode == DebounceModeActive {
		i.debounceTimer = time.AfterFunc(DebounceModeActive_Interval, i.rotate)
	}
	if i.debounceMode == DebounceModeAggressive {
		i.debounceTimer = time.AfterFunc(DebounceModeAggressive_Interval, i.rotate)
	}
}

func (i *stream) open() error {
	db, err := bolt.Open(fmt.Sprintf("%s.db", i.Path), 0600, nil)
	if err != nil {
		return err
	}
	i.handleDb = db
	return nil
}

func (i *stream) close() error {
	i.handleWg.Wait()
	if i.handleDb == nil {
		return nil
	}
	return i.handleDb.Close()
}

func (i *stream) rotate() {
	i.mx.Lock()
	go func() {
		defer i.mx.Unlock()

		i.debounceMx.Lock()
		defer i.debounceMx.Unlock()

		if i.debounceTimer != nil {
			i.debounceTimer.Stop()
			i.debounceTimer = nil
		} else {
			return
		}

		if err := i.close(); err != nil {
			fmt.Fprintf(os.Stderr, "error closing database: %s\n", err)
			return
		}
		p := fmt.Sprintf("%s-%d.db", i.Path, time.Now().Unix())
		if err := os.Rename(fmt.Sprintf("%s.db", i.Path), p); err != nil {
			fmt.Fprintf(os.Stderr, "error rotating file: %s\n", err)
		}
		if err := i.open(); err != nil {
			fmt.Fprintf(os.Stderr, "error opening database: %s\n", err)
			panic(err)
		}

		if i.fileC != nil {
			i.fileC <- p
		}
	}()
}

func (i *stream) Get() Handle {
	i.mx.RLock()
	defer i.mx.RUnlock()

	i.handleWg.Add(1)
	return &handle{i.handleDb, i}
}

func (i *stream) Put(_ Handle) {
	i.handleWg.Done()
}

func (i *stream) Close() error {
	i.rotate()
	i.mx.Lock()

	if i.fileC != nil {
		close(i.fileC)
	}
	time.Sleep(1 * time.Second)
	i.cancel()
	return i.close()
}
