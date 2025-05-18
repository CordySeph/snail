package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🐌 Snail CLI version:", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
