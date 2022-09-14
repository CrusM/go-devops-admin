package cmd

import (
	"go-devops-admin/database/migrations"
	"go-devops-admin/pkg/migrate"

	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "run not migrated migration",
	Run:   runUp,
}

var CmdMigrateDown = &cobra.Command{
	Use:   "down",
	Short: "Reverse the up command",
	Run:   runDown,
}

var CmdMigrateReset = &cobra.Command{
	Use:   "reset",
	Short: "Rollback all database migrations",
	Run:   runReset,
}

var CmdMigrateRefresh = &cobra.Command{
	Use:   "refresh",
	Short: "Reset and re-run all migrations",
	Run:   runRefresh,
}

var CmdMigrateFresh = &cobra.Command{
	Use:   "fresh",
	Short: "Drop all table and re-run all migrations",
	Run:   runFresh,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateDown,
		CmdMigrateReset,
		CmdMigrateRefresh,
		CmdMigrateFresh,
	)
}

func migrator() *migrate.Migrator {
	// 注册 database/migrations 下的所有文件
	migrations.Initialize()

	//
	return migrate.NewMigration()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().RollBack()
}

func runReset(cmd *cobra.Command, args []string) {
	migrator().Reset()
}

func runRefresh(cmd *cobra.Command, args []string) {
	migrator().Refresh()
}

func runFresh(cmd *cobra.Command, args []string) {
	migrator().Fresh()
}
