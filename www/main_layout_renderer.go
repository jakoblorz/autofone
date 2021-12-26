package www

import "github.com/gofiber/fiber/v2"

type RenderMainLayoutProps interface {
	GetTitle() string
	GetSlug() string
}

type MainLayoutProps struct {
	Title string
	Slug  string
}

func (p *MainLayoutProps) GetTitle() string {
	return p.Title
}

func (p *MainLayoutProps) GetSlug() string {
	return p.Slug
}

func RenderMainLayout(template string, c *fiber.Ctx, props RenderMainLayoutProps) error {
	p := fiber.Map{
		"Title": props.GetTitle(),
		"Slug":  props.GetSlug(),
	}
	if a, ok := props.(AdditionalPropsAppender); ok {
		p = a.AppendAdditionalProps(p)
	}
	return c.Render(template, p, "layouts/main")
}
