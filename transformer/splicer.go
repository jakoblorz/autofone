package transformer

import (
	"context"
	"sync"

	"github.com/jakoblorz/f1-metrics-transformer/pkg/step"
)

func Splicer(ctx context.Context, onPacketReceived func(pack interface{})) step.Step {
	pool := &sync.Pool{
		New: func() interface{} {
			r := &spliceRoutine{
				ctx:              ctx,
				in:               make(chan interface{}),
				onPacketReceived: onPacketReceived,
			}
			go r.process()
			return r
		},
	}
	return step.Intermediate(ctx, func(pack interface{}) interface{} {
		go func(pack interface{}) {
			s := pool.Get().(*spliceRoutine)
			s.in <- pack
			pool.Put(s)
		}(pack)
		return pack
	})
}

type spliceRoutine struct {
	ctx              context.Context
	in               chan interface{}
	onPacketReceived func(pack interface{})
}

func (m *spliceRoutine) process() {
	for {
		select {
		case <-m.ctx.Done():
			for {
				_, ok := <-m.in
				if !ok {
					return
				}
			}
		case msg := <-m.in:
			m.onPacketReceived(msg)
		}
	}
}
