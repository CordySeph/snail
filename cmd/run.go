package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [mode]",
	Short: "Run project (start | dev)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mode := args[0]
		switch mode {
		case "start":
			runStart()
		case "dev":
			runDev()
		default:
			fmt.Println("Usage: snail run [start|dev]")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runStart() {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	fmt.Println("🚀 Starting server (normal mode)...")
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error:", err)
	}
}

func runDev() {
	// ตรวจสอบว่า air ติดตั้งหรือยัง
	_, err := exec.LookPath("air")
	if err != nil {
		fmt.Println("❌ 'air' not found. Please install it with:")
		fmt.Println("go install github.com/cosmtrek/air@latest")
		return
	}

	cmd := exec.Command("air")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	fmt.Println("🔄 Starting server in dev mode (hot reload)...")
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error:", err)
	}
}
