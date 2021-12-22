package step

import "context"

type Step interface {
	context.Context

	ReadFrom(<-chan interface{})
	WriteTo(chan interface{})

	Out() <-chan interface{}

	// Starts the processing for this processor
	Process()

	// Then gets the next processor and returns the next processor,
	// so that Then(p1).Then(p2) etc is possible
	Then(Step) Step
}

type step struct {
	context.Context
	in  <-chan interface{}
	out chan interface{}

	processFunc func(in interface{}) interface{}
	emitFunc    func(out chan<- interface{})
}

func (p *step) ReadFrom(ch <-chan interface{}) {
	p.in = ch
}

func (p *step) WriteTo(ch chan interface{}) {
	p.out = ch
}

func (p *step) Out() <-chan interface{} {
	return p.out
}

func (p *step) Process() {
	defer close(p.out)

	if p.in == nil && p.emitFunc != nil {
		in := make(chan interface{}, 1)
		go p.emitFunc(in)
		p.in = in
	}

	for {
		select {
		case <-p.Done():
			for {
				_, ok := <-p.in
				if !ok {
					return
				}
			}
		case msg := <-p.in:
			if out := p.processFunc(msg); out != nil {
				p.out <- out
			}
		}
	}
}

func (p *step) Then(s Step) Step {
	s.ReadFrom(p.out)
	return s
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
