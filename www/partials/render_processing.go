package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/layouts"
)

type RenderProcessingSharedProps struct{}

func (r *RenderProcessingSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderProcessingPartialProps struct {
	layouts.PartialLayoutProps
	RenderProcessingSharedProps
}

func RenderProcessingPartial(c *fiber.Ctx) error {
	return layouts.RenderPartialLayout("partials/processing", c, &renderProcessingPartialProps{
		PartialLayoutProps: layouts.PartialLayoutProps{
			Slug: "processing",
		},
	})
}

type renderProcessingPageProps struct {
	layouts.MainLayoutProps
	RenderProcessingSharedProps
}

func RenderProcessingPage(c *fiber.Ctx) error {
	return layouts.RenderMainLayout("partials/processing", c, &renderProcessingPageProps{
		MainLayoutProps: layouts.MainLayoutProps{
			Title: "Processing - metrikx",
			Slug:  "processing",
		},
	})
}

var RenderProcessingHeader = renderHeader("processing")
