package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www"
)

type renderConfigurationPartialProps struct {
	www.PartialLayoutProps
}

func (r *renderConfigurationPartialProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

func RenderConfiguration(c *fiber.Ctx) error {
	return www.RenderPartialLayout("partials/configuration", c, &renderConfigurationPartialProps{
		PartialLayoutProps: www.PartialLayoutProps{
			Slug: "configuration",
		},
	})
}
