package main

import (
	"fmt"
	"gmm/lib/log"
	"gmm/lib/version"
	emoji "gmm/service.emoji"
	"os"

	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var logger = log.Package("entrypoint")

var rootCmd = &cobra.Command{
	Use:   "gmm",
	Short: "gmm monobinary",
}

func init() {
	rootCmd.AddCommand(emoji.CMD)
}

func main() {
	logger.Info(
		"starting gmm",
		zap.String("BuildTime", version.BuildTime),
		zap.String("CommitHash", version.CommitHash),
	)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
