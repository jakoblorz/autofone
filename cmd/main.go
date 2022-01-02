package cmd

import (
	"context"
	"expvar"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jakoblorz/metrikxd/modules"
	"github.com/jakoblorz/metrikxd/pipe"
	"github.com/jakoblorz/metrikxd/pkg/log"
	"github.com/jakoblorz/metrikxd/www"
	"github.com/jakoblorz/metrikxd/www/root"
	"github.com/spf13/cobra"
	"github.com/webview/webview"
	"github.com/zserge/metric"
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

			// // Random numbers always look nice on graphs
			expvar.Publish("random:gauge", metric.NewGauge("60s1s"))
			go func() {
				for range time.Tick(123 * time.Millisecond) {
					expvar.Get("random:gauge").(metric.Metric).Add(rand.Float64())
				}
			}()

			// Create a new engine by passing the template folder
			// and template extension using <engine>.New(dir, ext string)
			templates := html.New("./www", ".html")
			templates.AddFunc("path", path)
			templates.AddFunc("duration", duration)

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

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			r := modules.NewUDPPacketReader(ctx, "localhost", 20777, &pipe.PacketReaderOptions{})
			r.Mount(app)

			h := pipe.HandleEvents(ctx, pipe.EventHandler{})
			w := modules.NewHTTPPacketWriter(ctx, "api/udp/{{packetID}}")
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

			go app.Listen(":8080")

			view := webview.New(true)
			defer view.Destroy()
			view.SetTitle("metrix UI - F1 2021 UDP Utility")
			view.SetSize(1500, 1000, webview.HintNone)
			view.Navigate("http://localhost:8080")
			view.Run()
		},
	}
)

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

func path(samples []interface{}, keys ...string) []string {
	var min, max float64
	paths := make([]string, len(keys), len(keys))
	for i := 0; i < len(samples); i++ {
		s := samples[i].(map[string]interface{})
		for _, k := range keys {
			x := s[k].(float64)
			if i == 0 || x < min {
				min = x
			}
			if i == 0 || x > max {
				max = x
			}
		}
	}
	for i := 0; i < len(samples); i++ {
		s := samples[i].(map[string]interface{})
		for j, k := range keys {
			v := s[k].(float64)
			x := float64(i+1) / float64(len(samples))
			y := (v - min) / (max - min)
			if max == min {
				y = 0
			}
			if i == 0 {
				paths[j] = fmt.Sprintf("M%f %f", 0.0, (1-y)*18+1)
			}
			paths[j] += fmt.Sprintf(" L%f %f", x*100, (1-y)*18+1)
		}
	}
	return paths
}

func duration(samples []interface{}, n float64) string {
	n = n * float64(len(samples))
	if n < 60 {
		return fmt.Sprintf("%d sec", int(n))
	} else if n < 60*60 {
		return fmt.Sprintf("%d min", int(n/60))
	} else if n < 24*60*60 {
		return fmt.Sprintf("%d hrs", int(n/60/60))
	}
	return fmt.Sprintf("%d days", int(n/24/60/60))
}
