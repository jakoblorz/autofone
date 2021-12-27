package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www"
)

type RenderSendingSharedProps struct{}

func (r *RenderSendingSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderSendingPartialProps struct {
	www.PartialLayoutProps
	RenderSendingSharedProps
}

func RenderSendingPartial(c *fiber.Ctx) error {
	return www.RenderPartialLayout("partials/sending", c, &renderSendingPartialProps{
		PartialLayoutProps: www.PartialLayoutProps{
			Slug: "sending",
		},
	})
}

type renderSendingPageProps struct {
	www.MainLayoutProps
	RenderSendingSharedProps
}

func RenderSendingPage(c *fiber.Ctx) error {
	return www.RenderMainLayout("partials/sending", c, &renderSendingPageProps{
		MainLayoutProps: www.MainLayoutProps{
			Title: "Sending - metrikx",
			Slug:  "sending",
		},
	})
}
