package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jakoblorz/metrikxd/modules"
	"github.com/jakoblorz/metrikxd/pipe"
	"github.com/jakoblorz/metrikxd/pkg/log"
	"github.com/jakoblorz/metrikxd/www"
	"github.com/jakoblorz/metrikxd/www/root"
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

			app := fiber.New(fiber.Config{
				Views: templates,
			})

			app.Use(logger.New())

			app.Static("/", ".tailwindcss")

			app.Get("/", func(c *fiber.Ctx) error {
				return root.RenderIndexPage(c, "game-setup")
			})

			for _, p := range www.Pages {
				p.Mount(app)
			}

			app.Get("/chunk", adaptor.HTTPHandler(http.HandlerFunc(notifyChunk)))

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			r := modules.NewUDPPacketReader(ctx, "localhost", 20777, &pipe.PacketReaderOptions{})
			r.Mount(app)

			h := pipe.HandleEvents(ctx, pipe.EventHandler{})
			w := modules.NewHTTPPacketWriter(ctx, "localhost:8080/api/ingest/{{id}}")
			w.Mount(app)

			go r.Run()
			go h.Process()
			go w.Run()

			// r.Then(h).Then(w.Step())

			// conn, err := net.ListenUDP("udp", &net.UDPAddr{
			// 	IP:   net.ParseIP("localhost"),
			// 	Port: port,
			// })
			// if err != nil {
			// 	log.Printf("%+v", err)
			// 	return
			// }
			// defer conn.Close()

			// ctx, cancel := context.WithCancel(context.Background())
			// defer cancel()

			// r := pipe.ReadUDPPackets(ctx, conn, &pipe.PacketReaderOptions{})
			// h := pipe.HandleEvents(ctx, pipe.EventHandler{
			// 	OnLobbyInformationPacket: session.Instance.OnLobbyInfoDataReceived,
			// })
			// w := pipe.WritePacketToHTTP(ctx, to, pipe.JSONEncoding, pipe.StdoutResponseHandler)

			// r.Then(h).Then(w)

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
