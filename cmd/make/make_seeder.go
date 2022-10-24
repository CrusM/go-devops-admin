package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeSeeder = &cobra.Command{
	Use:   "seeder",
	Short: "create seeder file, example: make seeder user",
	Run:   runMakeSeeder,
	Args:  cobra.ExactArgs(1),
}

func runMakeSeeder(cmd *cobra.Command, args []string) {
	model := makeModelFormString(args[0])

	filepath := fmt.Sprintf("database/seeders/%s_seeder.go", model.TableName)

	createFileFromStub(filepath, "seeder/seeder", model)

}
