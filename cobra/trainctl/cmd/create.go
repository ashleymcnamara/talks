package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bketelsen/talks/cobra/trainctl/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new module skeleton",
	Long:  `Create a new module skeleton`,
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flag("name").Value.String() == "" {
			fmt.Println("missing required argument: --name [module name]")
			return
		}
		module := templates.NewModule(cmd,
			cmd.Flag("description").Value.String(),
			ProjectPath(),
		)
		err := createSubdirectories(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		createSlide(cmd, module)
		if err != nil {
			fmt.Println(err)
			return
		}

		createManifest(cmd, module)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().String("name", "", "Module name")
	createCmd.PersistentFlags().String("topic", "Go", "Module Topic {Go,Kubernetes}")
	createCmd.PersistentFlags().String("level", "beginner", "{beginner,intermediate,advanced,expert}")
	createCmd.PersistentFlags().String("description", "Module Description", "Module description")

}

func createSlide(cmd *cobra.Command, module templates.Module) error {
	name := cmd.Flag("name").Value.String() + ".slide"

	return writeTemplateToFile(ProjectPath(), name, templates.Slide, module)
}

func createManifest(cmd *cobra.Command, module templates.Module) error {
	name := cmd.Flag("name").Value.String() + ".json"

	js, err := json.Marshal(module)
	if err != nil {
		return errors.Wrap(err, "encoding module manifest")
	}
	return writeStringToFile(ProjectPath(), name, string(js))
}

func createSubdirectories(cmd *cobra.Command) error {
	path := filepath.Join(ProjectPath(), "src", cmd.Flag("name").Value.String())
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module src directory")
	}
	path = filepath.Join(ProjectPath(), "src", cmd.Flag("name").Value.String(), "exercises")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module exercises directory")
	}
	path = filepath.Join(ProjectPath(), "src", cmd.Flag("name").Value.String(), "solutions")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module solutions directory")
	}
	path = filepath.Join(ProjectPath(), "src", cmd.Flag("name").Value.String(), "demos")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "making module demos directory")
	}
	for _, dir := range subdirs {
		path := filepath.Join(ProjectPath(), cmd.Flag("name").Value.String(), dir)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return errors.Wrap(err, "making module slide directories")
		}
	}
	return nil
}
