package layouts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/rendering"
)

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
	if a, ok := props.(rendering.AdditionalPropsAppender); ok {
		p = a.AppendAdditionalProps(p)
	}
	return c.Render(template, p, "layouts/partial")
}
