package command

import (
	"github.com/spf13/cobra"

	gen_plugin "github.com/lab5e/gmqtt/cmd/gmqctl/command/gen-plugin"
)

// Gen is the command for code generator.
var Gen = &cobra.Command{
	Use:   "gen",
	Short: "Code generator",
}

func init() {
	Gen.AddCommand(gen_plugin.Command)
}
