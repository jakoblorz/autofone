package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www"
)

type RenderMonitoringSharedProps struct{}

func (r *RenderMonitoringSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	return m
}

type renderMonitoringPartialProps struct {
	www.PartialLayoutProps
	RenderMonitoringSharedProps
}

func RenderMonitoringPartial(c *fiber.Ctx) error {
	return www.RenderPartialLayout("partials/monitoring", c, &renderMonitoringPartialProps{
		PartialLayoutProps: www.PartialLayoutProps{
			Slug: "monitoring",
		},
	})
}

type renderMonitoringPageProps struct {
	www.MainLayoutProps
	RenderMonitoringSharedProps
}

func RenderMonitoringPage(c *fiber.Ctx) error {
	return www.RenderMainLayout("partials/monitoring", c, &renderMonitoringPageProps{
		MainLayoutProps: www.MainLayoutProps{
			Title: "Monitoring - metrikx",
			Slug:  "monitoring",
		},
	})
}
