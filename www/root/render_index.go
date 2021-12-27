package root

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/layouts"
)

type renderIndexPageProps struct {
	layouts.MainLayoutProps
	initialPageSlug string
}

func (r *renderIndexPageProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	m["InitialPageSlug"] = r.initialPageSlug
	return m
}

func RenderIndexPage(c *fiber.Ctx, initialPageSlug string) error {
	return layouts.RenderMainLayout("root/index", c, &renderIndexPageProps{
		MainLayoutProps: layouts.MainLayoutProps{
			Title: "metrikx",
			Slug:  initialPageSlug,
		},
		initialPageSlug: initialPageSlug,
	})
}
