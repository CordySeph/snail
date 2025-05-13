package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new project [name]",
	Short: "Create a new snail backend project",
	Args:  cobra.ExactArgs(2),
	// V.1
	// Run: func(cmd *cobra.Command, args []string) {
	// 	if args[0] != "project" {
	// 		fmt.Println("❌ Invalid command. Usage: snail new project <name>")
	// 		return
	// 	}
	// 	name := args[1]
	// 	createProject(name)
	// },

	// V.2
	Run: func(cmd *cobra.Command, args []string) {
		entity, name := args[0], args[1]

		switch entity {
		case "project":
			createProject(name)
		default:
			fmt.Printf("❌ Unsupported entity: %s\n", entity)
			fmt.Println("Usage: snail new project <name>")
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func createProject(name string) {
	paths := []string{
		filepath.Join(name, "modules"),
		filepath.Join(name, "config"),
		filepath.Join(name, "internal/logger"),
		filepath.Join(name, "internal/server"),
	}

	for _, p := range paths {
		if err := os.MkdirAll(p, 0755); err != nil {
			fmt.Println("❌ Failed to create directory:", p)
		}
	}

	writeProjectTemplate(filepath.Join(name, "main.go"), mainGoTemplate, name)
	writeProjectTemplate(filepath.Join(name, "go.mod"), goModTemplate, name)
	writeProjectTemplate(filepath.Join(name, "config/config.go"), configTemplate, name)
	writeProjectTemplate(filepath.Join(name, "internal/logger/logger.go"), loggerTemplate, name)
	writeProjectTemplate(filepath.Join(name, "internal/server/server.go"), serverTemplate, name)
	writeFile(filepath.Join(name, ".env"), "")
	writeFile(filepath.Join(name, ".air.toml"), airTomlContent)
	writeFile(filepath.Join(name, ".gitignore"), gitignoreContent)

	fmt.Printf("✅ Project '%s' created successfully!\n", name)
}

var gitignoreContent = `
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, build artifacts
*.test
*.out

# Output of the go coverage tool
*.coverprofile

# Dependency directories (if vendored)
vendor/

# IDE/editor folders
.idea/
.vscode/

# Air temp
tmp/
`

var airTomlContent = `
root = "."
tmp_dir = "tmp"
cmd = "go run main.go"
watch = ["."]
exclude_dir = ["tmp", "vendor"]
exclude_file = ["*.md", "*.toml"]
`

var mainGoTemplate = `package main

import (
	// "log"

	"{{.ProjectName}}/config"
	"{{.ProjectName}}/internal/logger"
	"{{.ProjectName}}/internal/server"
)

func main() {
	config.LoadEnv()
	log := logger.NewLogger()
	app := server.NewServer(log)

	log.Info("🚀 Server starting on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("❌ Server error: ", err)
	}
}
`

var goModTemplate = `module {{.ProjectName}}

go 1.21

require (
	github.com/gofiber/fiber/v2 v2.49.0
	github.com/joho/godotenv v1.5.1
	github.com/sirupsen/logrus v1.9.3
)
`

var configTemplate = `package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found")
	}
}
` + "\n"

var loggerTemplate = `package logger

import "github.com/sirupsen/logrus"

func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return log
}
`

var serverTemplate = `package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewServer(log *logrus.Logger) *fiber.App {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Infof("%s %s", c.Method(), c.Path())
		return c.Next()
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	return app
}
`
