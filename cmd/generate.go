package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	// "strings"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate module [name]",
	Short: "Generate a new module (handler, service, routes)",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] != "module" {
			fmt.Println("Usage: snail generate module <name>")
			return
		}
		name := args[1]
		generateModule(name)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateModule(name string) {
	base := filepath.Join("modules", name)
	os.MkdirAll(base, 0755)

	files := map[string]string{
		"handler.go": fmt.Sprintf("package %s\n\n// Handler for %s module", name, name),
		"service.go": fmt.Sprintf("package %s\n\n// Service for %s module", name, name),
		"routes.go":  fmt.Sprintf("package %s\n\n// Routes for %s module", name, name),
		"model.go":   fmt.Sprintf("package %s\n\n// Model for %s module", name, name),
	}

	for filename, content := range files {
		path := filepath.Join(base, filename)
		os.WriteFile(path, []byte(content), 0644)
	}
	fmt.Printf("✅ Created module: %s\n", name)
}
