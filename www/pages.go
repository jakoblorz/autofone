package www

import (
	"fmt"
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/partials"
)

var Pages = []Page{
	{"settings", partials.RenderSettingsPage, partials.RenderSettingsPartial, EmptySSEHandler},
	{"processing", partials.RenderProcessingPage, partials.RenderProcessingPartial, EmptySSEHandler},
	{"monitoring", partials.RenderMonitoringPage, partials.RenderMonitoringPartial, EmptySSEHandler},
	// {"sending", partials.RenderSendingPage, partials.RenderSendingPartial, EmptySSEHandler},
	{"workbench", nil, partials.RenderWorkbenchPartial, EmptySSEHandler},
}

func EmptySSEHandler(c *fiber.Ctx) error {
	return adaptor.HTTPHandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		rw.Header().Set("Content-Type", "text/event-stream")
		rw.Header().Set("Cache-Control", "no-cache")
		rw.Header().Set("Connection", "keep-alive")

		<-r.Context().Done()
	})(c)
}

type Page struct {
	Slug           string
	PageHandler    fiber.Handler
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
	if p.SSEHandler != nil {
		app.Get(fmt.Sprintf("/%s/sse", p.Slug), p.SSEHandler)
	}
}
