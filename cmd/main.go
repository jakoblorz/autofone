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
			app := fiber.New()
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
