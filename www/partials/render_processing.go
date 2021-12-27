package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www"
)

type RenderProcessingSharedProps struct{}

func (r *RenderProcessingSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderProcessingPartialProps struct {
	www.PartialLayoutProps
	RenderProcessingSharedProps
}

func RenderProcessingPartial(c *fiber.Ctx) error {
	return www.RenderPartialLayout("partials/processing", c, &renderProcessingPartialProps{
		PartialLayoutProps: www.PartialLayoutProps{
			Slug: "processing",
		},
	})
}

type renderProcessingPageProps struct {
	www.MainLayoutProps
	RenderProcessingSharedProps
}

func RenderProcessingPage(c *fiber.Ctx) error {
	return www.RenderMainLayout("partials/processing", c, &renderProcessingPageProps{
		MainLayoutProps: www.MainLayoutProps{
			Title: "Processing - metrikx",
			Slug:  "processing",
		},
	})
}
