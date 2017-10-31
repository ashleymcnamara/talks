package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// relatedCmd represents the related command
var relatedCmd = &cobra.Command{
	Use:   "related",
	Short: "Add related material to a module",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("related called")
		// UNIMPLEMENTED, the idea was to implement a command
		// that appends links to related course material to a
		// related.md file for the module.  Not sure yet if this
		// is useful, so not yet implemented.
	},
}

func init() {
	RootCmd.AddCommand(relatedCmd)
}
