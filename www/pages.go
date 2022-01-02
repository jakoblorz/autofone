package www

import (
	"bufio"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/partials"
)

var Pages = []Page{
	{"settings", partials.RenderSettingsPage, partials.RenderSettingsHeader, partials.RenderSettingsPartial, NotifyStatsChanged},
	{"processing", partials.RenderProcessingPage, partials.RenderProcessingHeader, partials.RenderProcessingPartial, NotifyStatsChanged},
	{"monitoring", partials.RenderMonitoringPage, partials.RenderMonitoringHeader, partials.RenderMonitoringPartial, NotifyStatsChanged},
}

func NotifyStatsChanged(c *fiber.Ctx) (err error) {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Headers", "Content-Type")
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		for {
			time.Sleep(1 * time.Second)

			fmt.Fprintf(w, "event: stats\ndata: %s\n\n", time.Now())

			if err = w.Flush(); err != nil {
				log.Print("client disconnected")
				return
			}
		}
	})
	return
}

type Page struct {
	Slug           string
	PageHandler    fiber.Handler
	HeaderHandler  fiber.Handler
	PartialHandler fiber.Handler
	SSEHandler     fiber.Handler
}

func (p *Page) Mount(app *fiber.App) {
	if p.PageHandler != nil {
		app.Get(fmt.Sprintf("/%s", p.Slug), p.PageHandler)
	}
	if p.PartialHandler != nil {
		app.Get(fmt.Sprintf("/p/%s", p.Slug), p.PartialHandler)
	}
	if p.HeaderHandler != nil {
		app.Get(fmt.Sprintf("/%s/header", p.Slug), p.HeaderHandler)
	}
	if p.SSEHandler != nil {
		app.Get(fmt.Sprintf("/%s/sse", p.Slug), p.SSEHandler)
	}
}
