package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about a module",
	Long:  `Get information about a module.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("name").Value.String() == "" {
			fmt.Println("--name parameter required")
			return
		}
		module, err := getManifest(cmd.Flag("name").Value.String())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(module)
	},
}

func init() {
	RootCmd.AddCommand(infoCmd)

	infoCmd.PersistentFlags().String("name", "", "Module name")
}
