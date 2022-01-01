package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/layouts"
)

type RenderSendingSharedProps struct {
	Host           string
	Port           int
	Encoding       string
	TemplateString string
}

func (r *RenderSendingSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderSendingPartialProps struct {
	layouts.PartialLayoutProps
	RenderSendingSharedProps
}

func RenderSendingPartial(c *fiber.Ctx, sharedProps RenderSendingSharedProps) error {
	return layouts.RenderPartialLayout("partials/sending", c, &renderSendingPartialProps{
		RenderSendingSharedProps: sharedProps,
		PartialLayoutProps: layouts.PartialLayoutProps{
			Slug: "sending",
		},
	})
}

type renderSendingPageProps struct {
	layouts.MainLayoutProps
	RenderSendingSharedProps
}

func RenderSendingPage(c *fiber.Ctx, sharedProps RenderSendingSharedProps) error {
	return layouts.RenderMainLayout("partials/sending", c, &renderSendingPageProps{
		RenderSendingSharedProps: sharedProps,
		MainLayoutProps: layouts.MainLayoutProps{
			Title: "Sending - metrikx",
			Slug:  "sending",
		},
	})
}
