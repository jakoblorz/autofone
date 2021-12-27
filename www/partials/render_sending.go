package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/layouts"
)

type RenderSendingSharedProps struct{}

func (r *RenderSendingSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderSendingPartialProps struct {
	layouts.PartialLayoutProps
	RenderSendingSharedProps
}

func RenderSendingPartial(c *fiber.Ctx) error {
	return layouts.RenderPartialLayout("partials/sending", c, &renderSendingPartialProps{
		PartialLayoutProps: layouts.PartialLayoutProps{
			Slug: "sending",
		},
	})
}

type renderSendingPageProps struct {
	layouts.MainLayoutProps
	RenderSendingSharedProps
}

func RenderSendingPage(c *fiber.Ctx) error {
	return layouts.RenderMainLayout("partials/sending", c, &renderSendingPageProps{
		MainLayoutProps: layouts.MainLayoutProps{
			Title: "Sending - metrikx",
			Slug:  "sending",
		},
	})
}
