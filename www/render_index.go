package www

import "github.com/gofiber/fiber/v2"

type renderIndexPageProps struct {
	MainLayoutProps
	initialPageSlug string
}

func (r *renderIndexPageProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	m["InitialPageSlug"] = r.initialPageSlug
	return m
}

func RenderIndexPage(c *fiber.Ctx, initialPageSlug string) error {
	return RenderMainLayout("index", c, &renderIndexPageProps{
		MainLayoutProps: MainLayoutProps{
			Title: "metrikx",
			Slug:  initialPageSlug,
		},
		initialPageSlug: initialPageSlug,
	})
}
