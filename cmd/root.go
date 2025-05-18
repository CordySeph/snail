package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "v0.1.3"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "snail",
	Short: "Snail CLI - A fast way to scaffold Go backend projects",
	Long: `Snail is a CLI tool to quickly scaffold Go backend projects with modular architecture.
It supports module generation, DB migration, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		showVersion, _ := cmd.Flags().GetBool("version")
		if showVersion {
			fmt.Println("🐌 Snail CLI version:", version)
			os.Exit(0)
		}
		_ = cmd.Help() // show help if no args provided
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Display CLI version")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
