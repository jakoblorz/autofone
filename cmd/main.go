package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/jakoblorz/autofone/pkg/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	verbose bool
	mac     string
	host    string

	rootCmd = &cobra.Command{
		Use: "autofone",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			config := zap.NewProductionConfig()
			if verbose {
				config = zap.NewDevelopmentConfig()
			}
			log.DefaultLogger, _ = config.Build()

			var err error
			host, err = os.Hostname()
			if err != nil {
				log.Printf("%+v", err)
				return
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Verbose output onto stdout")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}
