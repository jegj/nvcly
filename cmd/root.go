package cmd

import (
	"fmt"
	"os"
	"time"

	nvclyw "github.com/jegj/nvcly/widgets"
	"github.com/spf13/cobra"
)

var (
	version      = "dev"
	timeInterval time.Duration
)

const DEFAULT_TIME_INTERVAL = time.Second

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "nvcly",
	Short:         "Immersive terminal interface for managing nvidia-smi stats",
	Long:          "nvcly - Immersive terminal interface for managing nvidia-smi stats",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		nvclyw.InitiNvcly(timeInterval)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.PersistentFlags().DurationVarP(&timeInterval, "interval", "i", DEFAULT_TIME_INTERVAL, "Time interval to collects stats from nvidia-smi")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of nvcly",
	Long:  "Print the version number of nvcly",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("nvcly version %s\n", version)
	},
}
