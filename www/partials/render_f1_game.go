package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/layouts"
)

type RenderF1GameSharedProps struct {
	Host string
	Port int
}

func (r *RenderF1GameSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	m["Host"] = r.Host
	m["Port"] = r.Port
	return m
}

type renderF1GamePartialProps struct {
	layouts.PartialLayoutProps
	RenderF1GameSharedProps
}

func RenderF1GamePartial(c *fiber.Ctx, sharedProps RenderF1GameSharedProps) error {
	return layouts.RenderPartialLayout("partials/f1-game", c, &renderF1GamePartialProps{
		RenderF1GameSharedProps: sharedProps,
		PartialLayoutProps: layouts.PartialLayoutProps{
			Slug: "f1-game",
		},
	})
}

type renderF1GamePageProps struct {
	layouts.MainLayoutProps
	RenderF1GameSharedProps
}

func RenderF1GamePage(c *fiber.Ctx, sharedProps RenderF1GameSharedProps) error {
	return layouts.RenderMainLayout("partials/f1-game", c, &renderF1GamePageProps{
		RenderF1GameSharedProps: sharedProps,
		MainLayoutProps: layouts.MainLayoutProps{
			Title: "F1 Game - metrikx",
			Slug:  "f1-game",
		},
	})
}
