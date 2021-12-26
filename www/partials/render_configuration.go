package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www"
)

type RenderConfigurationSharedProps struct{}

func (r *RenderConfigurationSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderConfigurationPartialProps struct {
	www.PartialLayoutProps
	RenderConfigurationSharedProps
}

func RenderConfigurationPartial(c *fiber.Ctx) error {
	return www.RenderPartialLayout("partials/configuration", c, &renderConfigurationPartialProps{
		PartialLayoutProps: www.PartialLayoutProps{
			Slug: "configuration",
		},
	})
}

type renderConfigurationPageProps struct {
	www.MainLayoutProps
	RenderConfigurationSharedProps
}

func RenderConfigurationPage(c *fiber.Ctx) error {
	return www.RenderMainLayout("partials/configuration", c, &renderConfigurationPageProps{
		MainLayoutProps: www.MainLayoutProps{
			Title: "Configuration - metrikx",
			Slug:  "configuration",
		},
	})
}
