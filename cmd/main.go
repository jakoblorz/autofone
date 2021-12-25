package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/pipe"
	"github.com/jakoblorz/metrikxd/pkg/log"
	"github.com/jakoblorz/metrikxd/state_set/session"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	// In this example we use the html template engine
	"github.com/gofiber/template/html"
)

var (
	verbose bool
	rootCmd = &cobra.Command{
		Use: "metrikxd",
		PreRun: func(cmd *cobra.Command, args []string) {
			config := zap.NewProductionConfig()
			if verbose {
				config = zap.NewDevelopmentConfig()
			}
			log.DefaultLogger, _ = config.Build()
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Create a new engine by passing the template folder
			// and template extension using <engine>.New(dir, ext string)
			templates := html.New("./www", ".html")

			// Reload the templates on each render, good for development
			templates.Reload(true) // Optional. Default: false

			// Debug will print each template that is parsed, good for debugging
			templates.Debug(true) // Optional. Default: false

			// Layout defines the variable name that is used to yield templates within layouts
			templates.Layout("embed") // Optional. Default: "embed"

			app := fiber.New(fiber.Config{
				Views: templates,
			})

			app.Static("/", ".tailwindcss")

			// To render a template, you can call the ctx.Render function
			// Render(tmpl string, values interface{}, layout ...string)
			app.Get("/", func(c *fiber.Ctx) error {
				return c.Render("index", fiber.Map{
					"Title": "Hello, World!",
				}, "layouts/main")
			})

			app.Get("/chunk", adaptor.HTTPHandler(http.HandlerFunc(notifyChunk)))

			conn, err := net.ListenUDP("udp", &net.UDPAddr{
				IP:   net.ParseIP("localhost"),
				Port: port,
			})
			if err != nil {
				log.Printf("%+v", err)
				return
			}
			defer conn.Close()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			r := pipe.ReadUDPPackets(ctx, conn, &pipe.PacketReaderOptions{})
			h := pipe.HandleEvents(ctx, pipe.EventHandler{
				OnLobbyInformationPacket: session.Instance.OnLobbyInfoDataReceived,
			})
			w := pipe.WritePacketToHTTP(ctx, to, pipe.JSONEncoding, pipe.StdoutResponseHandler)

			r.Then(h).Then(w)

			app.Listen(":8080")
		},
	}
)

func notifyChunk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

func verbosef(format string, args ...interface{}) {
	if verbose {
		log.Printf(format, args...)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Verbose output onto stdout")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
