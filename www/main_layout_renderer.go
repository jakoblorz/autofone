package www

import "github.com/gofiber/fiber/v2"

type RenderMainLayoutProps interface {
	GetTitle() string
}

type MainLayoutProps struct {
	Title string
}

func (p *MainLayoutProps) GetTitle() string {
	return p.Title
}

func RenderMainLayout(template string, c *fiber.Ctx, props RenderMainLayoutProps) error {
	p := fiber.Map{
		"Title": props.GetTitle(),
	}
	if a, ok := props.(AdditionalPropsAppender); ok {
		p = a.AppendAdditionalProps(p)
	}
	return c.Render(template, p, "layout/main")
}
