package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www"
)

type RenderF1GameSharedProps struct{}

func (r *RenderF1GameSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderF1GamePartialProps struct {
	www.PartialLayoutProps
	RenderF1GameSharedProps
}

func RenderF1GamePartial(c *fiber.Ctx) error {
	return www.RenderPartialLayout("partials/f1-game", c, &renderF1GamePartialProps{
		PartialLayoutProps: www.PartialLayoutProps{
			Slug: "f1-game",
		},
	})
}

type renderF1GamePageProps struct {
	www.MainLayoutProps
	RenderF1GameSharedProps
}

func RenderF1GamePage(c *fiber.Ctx) error {
	return www.RenderMainLayout("partials/f1-game", c, &renderF1GamePageProps{
		MainLayoutProps: www.MainLayoutProps{
			Title: "F1 Game - metrikx",
			Slug:  "f1-game",
		},
	})
}
