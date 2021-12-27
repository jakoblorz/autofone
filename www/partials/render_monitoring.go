package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/layouts"
)

type RenderMonitoringSharedProps struct{}

func (r *RenderMonitoringSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderMonitoringPartialProps struct {
	layouts.PartialLayoutProps
	RenderMonitoringSharedProps
}

func RenderMonitoringPartial(c *fiber.Ctx) error {
	return layouts.RenderPartialLayout("partials/monitoring", c, &renderMonitoringPartialProps{
		PartialLayoutProps: layouts.PartialLayoutProps{
			Slug: "monitoring",
		},
	})
}

type renderMonitoringPageProps struct {
	layouts.MainLayoutProps
	RenderMonitoringSharedProps
}

func RenderMonitoringPage(c *fiber.Ctx) error {
	return layouts.RenderMainLayout("partials/monitoring", c, &renderMonitoringPageProps{
		MainLayoutProps: layouts.MainLayoutProps{
			Title: "Monitoring - metrikx",
			Slug:  "monitoring",
		},
	})
}
