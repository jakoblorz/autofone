package modules

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/pipe"
	"github.com/jakoblorz/metrikxd/pkg/log"
	"github.com/jakoblorz/metrikxd/pkg/step"
	"github.com/jakoblorz/metrikxd/www"
	"github.com/jakoblorz/metrikxd/www/partials"
)

type WriteHTTPPackets struct {
	context.Context
	step.Step
	www.Page

	toTemplate toTemplateRenderer
	toEncoding pipe.HTTPEncoding

	responseHandler pipe.HTTPResponseHandler

	stepCtx    context.Context
	stepCancel context.CancelFunc

	applyAfterWith   step.Outer
	applyAfterWithMx sync.Locker
}

func NewHTTPPacketWriter(ctx context.Context, initialTemplateString string) *WriteHTTPPackets {
	w := &WriteHTTPPackets{
		Context:          ctx,
		toTemplate:       toTemplateRenderer{"matri-x.de", 0, initialTemplateString, initialTemplateString},
		toEncoding:       pipe.JSONEncoding,
		responseHandler:  pipe.StdoutResponseHandler,
		applyAfterWithMx: new(sync.Mutex),
	}
	w.Page = www.Page{"sending", w.renderSendingPage, partials.RenderSendingHeader, w.renderSendingPartial, www.NotifyStatsChanged}
	return w
}

func (w *WriteHTTPPackets) getSharedProps() partials.RenderSendingSharedProps {
	return partials.RenderSendingSharedProps{
		Host:           w.toTemplate.Host,
		Port:           w.toTemplate.Port,
		Encoding:       string(w.toEncoding),
		TemplateString: w.toTemplate.toTemplateString,
	}
}

func (w *WriteHTTPPackets) renderSendingPage(c *fiber.Ctx) error {
	return partials.RenderSendingPage(c, w.getSharedProps())
}

func (w *WriteHTTPPackets) renderSendingPartial(c *fiber.Ctx) error {
	return partials.RenderSendingPartial(c, w.getSharedProps())
}

func (w *WriteHTTPPackets) Mount(app *fiber.App) {
	w.Page.Mount(app)
	app.Post(fmt.Sprintf("/%s", w.Page.Slug), w.updateHTTPWriter)
}

type UpdateHTTPWriterRequest struct {
	Host           string `form:"host"`
	Port           int    `form:"port"`
	Encoding       string `form:"encoding"`
	TemplateString string `form:"template_string"`
}

func (w *WriteHTTPPackets) updateHTTPWriter(c *fiber.Ctx) error {
	d := new(UpdateHTTPWriterRequest)
	if err := c.BodyParser(d); err != nil {
		log.Printf("%+v", err)
		return c.Redirect(w.Page.Slug)
	}

	s := fmt.Sprintf("https://%s", d.Host)
	if d.Port != 0 {
		s = fmt.Sprintf("%s:%d/", s, d.Port)
	} else {
		s = fmt.Sprintf("%s/", s)
	}
	if d.TemplateString != "" {
		var t string
		if strings.HasPrefix(d.TemplateString, "/") {
			t = strings.Replace(d.TemplateString, "/", "", 1)
		} else {
			t = d.TemplateString
		}
		s = fmt.Sprintf("%s%s", s, t)
	}
	if s != w.toTemplate.toTemplateString {
		w.setState(func() error {
			w.toTemplate = toTemplateRenderer{d.Host, d.Port, d.TemplateString, s}
			return nil
		})
	}
	return partials.RenderSendingPage(c, w.getSharedProps())
}

func (w *WriteHTTPPackets) setState(u func() error) error {
	defer w.stepCancel()
	return u()
}

func (w *WriteHTTPPackets) After(o step.Outer) {
	w.applyAfterWithMx.Lock()
	defer w.applyAfterWithMx.Unlock()

	w.applyAfterWith = o
	if w.stepCtx != nil {
		w.stepCancel()
	}
}

func (w *WriteHTTPPackets) Run() {
	for {
		select {
		case <-w.Context.Done():
			return
		default:
			func() {
				w.applyAfterWithMx.Lock()
				w.stepCtx, w.stepCancel = context.WithCancel(w.Context)
				w.Step = pipe.WritePacketToHTTP(w.stepCtx, &w.toTemplate, w.toEncoding, w.responseHandler)
				if w.applyAfterWith != nil {
					w.Step.After(w.applyAfterWith)
				}
				w.applyAfterWithMx.Unlock()

				w.Step.Process()
			}()
		}
	}
}

type toTemplateRenderer struct {
	Host           string
	Port           int
	TemplateString string

	toTemplateString string
}

func (t *toTemplateRenderer) String() string {
	return t.toTemplateString
}
