package step

import (
	"context"
	"time"
)

type Outer interface {
	Out() chan interface{}
}

type Inner interface {
	In() chan interface{}
}

type Step interface {
	Outer
	Inner

	ReadFrom(chan interface{})
	WriteTo(chan interface{})

	// Starts the processing for this processor
	Process()

	// Then gets the next processor and returns the next processor,
	// so that Then(p1).Then(p2) etc is possible
	After(Outer)
}

type step struct {
	context.Context

	in   chan interface{}
	prev Outer

	out chan interface{}

	processFunc func(in interface{}) interface{}
	emitFunc    func(out chan<- interface{})
}

func (p *step) ReadFrom(ch chan interface{}) {
	p.in = ch
}

func (p *step) WriteTo(ch chan interface{}) {
	p.out = ch
}

func (p *step) In() chan interface{} {
	return p.in
}

func (p *step) Out() chan interface{} {
	return p.out
}

func (p *step) Process() {

	if p.in == nil && p.prev == nil && p.emitFunc != nil {
		in := make(chan interface{}, 1)
		go p.emitFunc(in)
		p.in = in
	}

	for {
		if p.prev != nil {
			if in := p.prev.Out(); in != nil {
				select {
				case <-p.Done():
					for {
						_, ok := <-in
						if !ok {
							return
						}
					}
				case msg, ok := <-in:
					if !ok {
						return
					}

					if out := p.processFunc(msg); out != nil {
						p.out <- out
					}
				}
			} else {
				<-time.After(100 * time.Millisecond)
			}
		} else {
			select {
			case <-p.Done():
				for {
					_, ok := <-p.in
					if !ok {
						return
					}
				}
			case msg, ok := <-p.in:
				if !ok {
					return
				}

				if out := p.processFunc(msg); out != nil {
					p.out <- out
				}
			}
		}
	}
}

func (p *step) After(s Outer) {
	p.prev = s
	p.in = nil
}

func New(ctx context.Context, processFunc func(in interface{}) interface{}, emitFunc func(out chan<- interface{})) Step {
	if ctx == nil {
		ctx = context.Background()
	}
	return &step{
		Context:     ctx,
		out:         make(chan interface{}, 1),
		processFunc: processFunc,
		emitFunc:    emitFunc,
	}
}

func Emitter(ctx context.Context, emitFunc func(out chan<- interface{})) Step {
	return New(ctx, func(in interface{}) interface{} {
		return in
	}, emitFunc)
}

func Intermediate(ctx context.Context, processFunc func(interface{}) interface{}) Step {
	return New(ctx, processFunc, nil)
}
