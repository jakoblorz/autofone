package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/layouts"
)

type RenderWorkbenchSharedProps struct{}

func (r *RenderWorkbenchSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderWorkbenchPartialProps struct {
	layouts.PartialLayoutProps
	RenderWorkbenchSharedProps
}

func RenderWorkbenchPartial(c *fiber.Ctx) error {
	return c.Render("partials/workbench", &renderWorkbenchPartialProps{
		PartialLayoutProps: layouts.PartialLayoutProps{
			Slug: string(c.Request().URI().QueryArgs().Peek("slug")),
		},
	})
}
