package make

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create policy file,example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1),
}

func runMakePolicy(cmd *cobra.Command, args []string) {
	model := makeModelFormString(args[0])

	dir := fmt.Sprintf("app/%s/policies/", model.PackageName)
	os.MkdirAll(dir, os.ModePerm)
	filepath := fmt.Sprintf("%s/policies.go", dir)

	createFileFromStub(filepath, "policies/policies", model)
}
