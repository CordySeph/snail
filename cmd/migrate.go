package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Manage database migrations",
}

var migrateInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize migrations directory",
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.MkdirAll("migrations", 0755); err != nil {
			fmt.Println("❌ Failed to create migrations directory:", err)
			return
		}
		fmt.Println("✅ Migrations directory created.")
	},
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new migration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := strings.ToLower(strings.ReplaceAll(args[0], " ", "_"))
		timestamp := time.Now().Format("20060102150405")
		up := filepath.Join("migrations", fmt.Sprintf("%s_%s.up.sql", timestamp, name))
		down := filepath.Join("migrations", fmt.Sprintf("%s_%s.down.sql", timestamp, name))

		os.WriteFile(up, []byte("-- Write your UP migration here\n"), 0644)
		os.WriteFile(down, []byte("-- Write your DOWN migration here\n"), 0644)

		fmt.Println("✅ Migration files created:")
		fmt.Println("  ", up)
		fmt.Println("  ", down)
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all up migrations",
	Run: func(cmd *cobra.Command, args []string) {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			fmt.Println("❌ DATABASE_URL not set in .env")
			return
		}
		exec.Command("migrate", "-database", dsn, "-path", "migrations", "up").Run()
		fmt.Println("✅ Migrations applied.")
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback the last migration",
	Run: func(cmd *cobra.Command, args []string) {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			fmt.Println("❌ DATABASE_URL not set in .env")
			return
		}
		exec.Command("migrate", "-database", dsn, "-path", "migrations", "down", "1").Run()
		fmt.Println("✅ Rolled back the last migration.")
	},
}

func init() {
	migrateCmd.AddCommand(migrateInitCmd)
	migrateCmd.AddCommand(migrateCreateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrateCmd)
}
