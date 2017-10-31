package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// removeCmd represents the create command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a new module skeleton",
	Long:  `Removes the module directory and all of its contents, as well as the directory`,
	Run: func(cmd *cobra.Command, args []string) {
		really, _ := cmd.PersistentFlags().GetBool("really")
		if really {
			err := removeSubdirectories(cmd)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}
		fmt.Printf("Cowardly refusing to remove module %s without --really flag", cmd.Flag("name").Value.String())
		return
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	removeCmd.PersistentFlags().String("name", "", "Module name")

	removeCmd.PersistentFlags().Bool("really", false, "Really destroy this module? It's permanent!")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func removeSubdirectories(cmd *cobra.Command) error {

	path := filepath.Join(ProjectPath(), "src", cmd.Flag("name").Value.String())
	err := os.RemoveAll(path)
	if err != nil {
		return errors.Wrap(err, "removing module src directory")
	}

	name := cmd.Flag("name").Value.String()
	path = getPath(ProjectPath(), name)
	fmt.Println("removing module subdirectory:", path)
	err = os.RemoveAll(path)
	if err != nil {
		return errors.Wrap(err, "removing module slide")
	}

	name = cmd.Flag("name").Value.String()
	path = getPath(ProjectPath(), name+".slide")
	fmt.Println("removing slide:", path)
	err = os.RemoveAll(path)
	if err != nil {
		return errors.Wrap(err, "removing module slide")
	}

	path = getPath(ProjectPath(), name+".json")

	fmt.Println("removing manifest:", path)
	err = os.RemoveAll(path)
	if err != nil {
		return errors.Wrap(err, "removing module manifest")
	}
	return nil

}
