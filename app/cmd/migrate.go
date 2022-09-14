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

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
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
