package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "trainctl",
	Short: "An application to manage training material",
	Long:  `trainctl is an opinionated tool to help manage training material.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.trainctl.yaml)")

	RootCmd.PersistentFlags().StringP("author", "a", "", "course or module author name")
	RootCmd.PersistentFlags().StringP("email", "e", "", "course or module author email")
	RootCmd.PersistentFlags().StringP("twitter", "t", "", "course or module author twitter handle")
	RootCmd.PersistentFlags().StringP("moduledir", "m", "", "module location path")
	RootCmd.PersistentFlags().StringP("coursedir", "c", "", "course location path")

	viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("email", RootCmd.PersistentFlags().Lookup("email"))
	viper.BindPFlag("twitter", RootCmd.PersistentFlags().Lookup("twitter"))
	viper.BindPFlag("moduledir", RootCmd.PersistentFlags().Lookup("moduledir"))
	viper.BindPFlag("coursedir", RootCmd.PersistentFlags().Lookup("coursedir"))

	viper.SetDefault("author", "NAME HERE")
	viper.SetDefault("email", "you@email.com")
	viper.SetDefault("twitter", "you")
	viper.SetDefault("moduledir", "~/courses")
	viper.SetDefault("coursedir", "~/modules")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName(".trainctl") // name of config file (without extension)
	viper.AddConfigPath("$HOME")     // adding home directory as first search path
	viper.AutomaticEnv()             // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	if err != nil {
		fmt.Println("config file error: ", err)
	}

}
