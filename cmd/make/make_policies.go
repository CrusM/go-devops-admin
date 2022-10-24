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

	os.MkdirAll("app/policies", os.ModePerm)

	filepath := fmt.Sprintf("app/policies/%s_policies.go", model.PackageName)

	createFileFromStub(filepath, "policies/policies", model)
}
