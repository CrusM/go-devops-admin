package make

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "create model's factory file, example: make factory user",
	Run:   runMakeFactory,
	Args:  cobra.ExactArgs(1),
}

func runMakeFactory(cmd *cobra.Command, args []string) {
	model := makeModelFormString(args[0])

	filepath := fmt.Sprintf("database/factories/%s_factory.go", model.PackageName)

	createFileFromStub(filepath, "factory/factory", model)
}
