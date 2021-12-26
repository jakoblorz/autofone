package www

import "github.com/gofiber/fiber/v2"

type RenderPartialLayoutProps interface {
	GetSlug() string
}

type PartialLayoutProps struct {
	Slug string
}

func (p *PartialLayoutProps) GetSlug() string {
	return p.Slug
}

func RenderPartialLayout(template string, c *fiber.Ctx, props RenderPartialLayoutProps) error {
	p := fiber.Map{
		"Slug": props.GetSlug(),
	}
	if a, ok := props.(AdditionalPropsAppender); ok {
		p = a.AppendAdditionalProps(p)
	}
	return c.Render(template, p, "layouts/partial")
}
