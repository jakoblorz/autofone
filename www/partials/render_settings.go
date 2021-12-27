package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/layouts"
)

type RenderSettingsSharedProps struct{}

func (r *RenderSettingsSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderSettingsPartialProps struct {
	layouts.PartialLayoutProps
	RenderSettingsSharedProps
}

func RenderSettingsPartial(c *fiber.Ctx) error {
	return layouts.RenderPartialLayout("partials/settings", c, &renderSettingsPartialProps{
		PartialLayoutProps: layouts.PartialLayoutProps{
			Slug: "settings",
		},
	})
}

type renderSettingsPageProps struct {
	layouts.MainLayoutProps
	RenderSettingsSharedProps
}

func RenderSettingsPage(c *fiber.Ctx) error {
	return layouts.RenderMainLayout("partials/settings", c, &renderSettingsPageProps{
		MainLayoutProps: layouts.MainLayoutProps{
			Title: "Settings - metrikx",
			Slug:  "settings",
		},
	})
}
