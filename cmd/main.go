package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool
	rootCmd = &cobra.Command{
		Use: "f1-metrics-transformer",
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
