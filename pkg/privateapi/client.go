package privateapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	clientPool = &sync.Pool{
		New: func() interface{} {
			return &http.Client{}
		},
	}
)

type Client interface {
	io.Closer
	Devices() *DevicesClient
	Snapshots() *SnapshotsClient
}

func New(token string, baseURL string) Client {
	ctx, cancel := context.WithCancel(context.Background())
	c := &client{
		Mutex:   sync.Mutex{},
		token:   token,
		baseURL: baseURL,
		ctx:     ctx,
		cancel:  cancel,
	}
	go c.syncState()
	go c.syncHealth()
	return c
}

type client struct {
	sync.Mutex
	token   string
	baseURL string

	ctx    context.Context
	cancel context.CancelFunc
}

func (c *client) syncState() {
	errCounter := 0
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(5 * time.Minute):
			t, err := c.Devices().RefreshToken()
			if err != nil {
				errCounter++
				if errCounter > 10 {
					panic(err)
				} else {
					log.Printf("error refreshing token: %s", err)
				}
				continue
			}
			if errCounter > 0 {
				log.Printf("token refreshed after %d errors", errCounter)
			}
			errCounter = 0
			c.Lock()
			c.token = t.Token
			c.Unlock()
		}
	}
}

func (c *client) syncHealth() {
	errCounter := 0
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			err := c.Snapshots().Health()
			if err != nil {
				errCounter++
				if errCounter > 10 {
					log.Printf("error streaming health status, panicking: %s", err)
					panic(err)
				} else {
					log.Printf("error streaming health status: %s", err)
				}
				continue
			}
		}
	}
}

func (c *client) Close() error {
	c.cancel()
	return nil
}

func (c *client) Devices() *DevicesClient {
	return &DevicesClient{c}
}

func (c *client) Snapshots() *SnapshotsClient {
	return &SnapshotsClient{c}
}

func (c *client) JSON(method, url string, body interface{}, target interface{}) error {
	c.Lock()
	var (
		token   = c.token
		baseURL = c.baseURL
	)
	c.Unlock()

	var (
		err error
		req *http.Request
	)
	if body != nil {
		var r []byte
		r, err = json.Marshal(r)
		if err != nil {
			return err
		}
		req, err = http.NewRequest(method, fmt.Sprintf("%s%s", baseURL, url), bytes.NewBuffer(r))
	} else {
		req, err = http.NewRequest(method, fmt.Sprintf("%s%s", baseURL, url), nil)
	}
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	client := clientPool.Get().(*http.Client)
	res, err := client.Do(req)
	clientPool.Put(client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
		return json.NewDecoder(res.Body).Decode(target)
	}
	return fmt.Errorf("unexpected status code %d", res.StatusCode)
}
