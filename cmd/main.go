package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jakoblorz/metrikxd/pkg/log"
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
