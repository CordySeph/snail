package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate module [name]",
	Short: "Generate a new module with handler, service, and route",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		entity, name := args[0], args[1]
		switch entity {
		case "module":
			generateModule(name)
		default:
			fmt.Printf("❌ Unsupported entity: %s\n", entity)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateModule(name string) {
	modulePath := filepath.Join("modules", name)
	if err := os.MkdirAll(modulePath, 0755); err != nil {
		fmt.Println("❌ Failed to create module directory:", err)
		return
	}

	writeModuleTemplate(filepath.Join(modulePath, "handler.go"), handlerTemplate, name)
	writeModuleTemplate(filepath.Join(modulePath, "service.go"), serviceTemplate, name)
	writeModuleTemplate(filepath.Join(modulePath, "route.go"), routeTemplate, name)

	fmt.Printf("✅ Module '%s' generated successfully!\n", name)
}

var handlerTemplate = `package {{.}}

import "github.com/gofiber/fiber/v2"

func Get{{title .}}(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "GET {{.}}",
	})
}
`

var serviceTemplate = `package {{.}}

func Get{{title .}}Data() string {
	return "data from {{.}} service"
}
`

var routeTemplate = `package {{.}}

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	r := app.Group("/{{.}}")
	r.Get("/", Get{{title .}})
}
`

func upperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}
