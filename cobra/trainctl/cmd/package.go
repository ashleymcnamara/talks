package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "package a course into a downloadable zip file",
	Long:  `Package a course into a downloadable zip file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("shortname").Value.String() == "" {
			fmt.Println("--shortname parameter required")
			return
		}
		//zip -r thing.zip thing
		courseDir := filepath.Join(viper.GetString("coursedir"), cmd.Flag("shortname").Value.String())
		exists, err := dirExists(courseDir)
		if err != nil {
			fmt.Println(err, "package directory check")
			return
		}
		if exists {
			zipName := courseDir + ".zip"

			cmd := exec.Command("zip", "-r", zipName, "./"+cmd.Flag("shortname").Value.String())
			cmd.Dir = viper.GetString("coursedir")
			cmd.Stdout = os.Stdout

			if err := cmd.Run(); err != nil {
				fmt.Println(err, "create zip")
				return
			}
			return
		}

		fmt.Println("unknown course")
		return

	},
}

func init() {
	RootCmd.AddCommand(packageCmd)

	packageCmd.PersistentFlags().String("shortname", "", "Course short name")
}
