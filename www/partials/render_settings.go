package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www"
)

type RenderSettingsSharedProps struct{}

func (r *RenderSettingsSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderSettingsPartialProps struct {
	www.PartialLayoutProps
	RenderSettingsSharedProps
}

func RenderSettingsPartial(c *fiber.Ctx) error {
	return www.RenderPartialLayout("partials/settings", c, &renderSettingsPartialProps{
		PartialLayoutProps: www.PartialLayoutProps{
			Slug: "settings",
		},
	})
}

type renderSettingsPageProps struct {
	www.MainLayoutProps
	RenderSettingsSharedProps
}

func RenderSettingsPage(c *fiber.Ctx) error {
	return www.RenderMainLayout("partials/settings", c, &renderSettingsPageProps{
		MainLayoutProps: www.MainLayoutProps{
			Title: "Settings - metrikx",
			Slug:  "settings",
		},
	})
}
