package cmd

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/jakoblorz/autofone/pkg/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

var (
	verbose bool
	mac     string
	host    string

	credentialJSON string

	storageClient *storage.Client
	storageBucket = "streamdb-content"
	projectID     = "autofone-355408"
	sessionID     string

	rootCmd = &cobra.Command{
		Use: "autofone",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			config := zap.NewProductionConfig()
			if verbose {
				config = zap.NewDevelopmentConfig()
			}
			log.DefaultLogger, _ = config.Build()

			credentialData, err := base64.StdEncoding.DecodeString(credentialJSON)
			if err != nil {
				log.Print(err)
				return
			}

			macAddresses, err := getMacAddr()
			if err != nil {
				log.Print(err)
				return
			}
			if len(macAddresses) == 0 {
				log.Print(errors.New("no mac address found"))
				return
			}
			mac = strings.ReplaceAll(macAddresses[0], ":", "")
			log.Printf("Using MAC Address %s for identification", mac)

			storageClient, err = storage.NewClient(context.Background(), option.WithCredentialsJSON(credentialData))
			if err != nil {
				log.Print(err)
				return
			}

			host, err = os.Hostname()
			if err != nil {
				log.Printf("%+v", err)
				return
			}
			sessionID = fmt.Sprintf("%s-%d", host, int64(float64(time.Now().UnixNano())*rand.New(rand.NewSource(time.Now().UnixNano())).Float64()))

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
