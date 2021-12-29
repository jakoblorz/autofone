package modules

import (
	"context"

	"github.com/jakoblorz/metrikxd/pipe"
	"github.com/jakoblorz/metrikxd/pkg/step"
	"github.com/jakoblorz/metrikxd/www"
	"github.com/jakoblorz/metrikxd/www/partials"
)

type WriteHTTPPackets struct {
	context.Context
	www.Page

	toTemplate toTemplateRenderer
	toEncoding pipe.HTTPEncoding

	responseHandler pipe.HTTPResponseHandler

	step step.Step

	applyThenWith step.Step
}

func NewHTTPPacketWriter(ctx context.Context, initialTemplateString string) *WriteHTTPPackets {
	w := &WriteHTTPPackets{
		Context:         ctx,
		toTemplate:      toTemplateRenderer{initialTemplateString},
		toEncoding:      pipe.JSONEncoding,
		responseHandler: pipe.StdoutResponseHandler,
	}
	w.Page = www.Page{"sending", partials.RenderSendingPage, partials.RenderSendingPartial, www.EmptySSEHandler}
	return w
}

func (w *WriteHTTPPackets) Run() {
	for {
		select {
		case <-w.Context.Done():
			return
		default:
			func() {
				w.step = pipe.WritePacketToHTTP(w.Context, &w.toTemplate, w.toEncoding, w.responseHandler)
				if w.applyThenWith != nil {
					w.step.Then(w.applyThenWith)
				}

				w.step.Process()
			}()
		}
	}
}

func (r *WriteHTTPPackets) Step() step.Step {
	return &stepPatcher{
		Step: r.step,
		onThen: func(st step.Step) step.Step {
			r.applyThenWith = st
			return st
		},
	}
}

type toTemplateRenderer struct {
	toTemplateString string
}

func (t *toTemplateRenderer) String() string {
	return t.toTemplateString
}

type stepPatcher struct {
	step.Step
	onThen func(step.Step) step.Step
}

func (s *stepPatcher) Then(st step.Step) step.Step {
	s.onThen(st)
	return st
}
