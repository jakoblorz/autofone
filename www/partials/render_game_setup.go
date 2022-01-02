package partials

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/www/layouts"
)

type SinglePacketFilter struct {
	Id      string
	Name    string
	Value   bool
	Content string
}

type RenderGameSetupSharedProps struct {
	Host    string
	Port    int
	Packets []SinglePacketFilter
}

func (r *RenderGameSetupSharedProps) AppendAdditionalProps(m fiber.Map) fiber.Map {
	m["Host"] = r.Host
	m["Port"] = r.Port
	m["Packets"] = r.Packets
	return m
}

type renderGameSetupPartialProps struct {
	layouts.PartialLayoutProps
	RenderGameSetupSharedProps
}

func RenderGameSetupPartial(c *fiber.Ctx, sharedProps RenderGameSetupSharedProps) error {
	return layouts.RenderPartialLayout("partials/game-setup", c, &renderGameSetupPartialProps{
		RenderGameSetupSharedProps: sharedProps,
		PartialLayoutProps: layouts.PartialLayoutProps{
			Slug: "game-setup",
		},
	})
}

type renderGameSetupPageProps struct {
	layouts.MainLayoutProps
	RenderGameSetupSharedProps
}

func RenderGameSetupPage(c *fiber.Ctx, sharedProps RenderGameSetupSharedProps) error {
	return layouts.RenderMainLayout("partials/game-setup", c, &renderGameSetupPageProps{
		RenderGameSetupSharedProps: sharedProps,
		MainLayoutProps: layouts.MainLayoutProps{
			Title: "Game Setup - metrikx",
			Slug:  "game-setup",
		},
	})
}

var RenderGameSetupHeader = renderHeader("game-setup")
