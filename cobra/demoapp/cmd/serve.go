// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var directory string
var port string
var address string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve delivers HTML from a named directory",
	Long: `serve delivers HTML from a directory which can
be specified with a parameter, or the current directory by default.`,
	Run: func(cmd *cobra.Command, args []string) {
		if directory == "" {
			var err error
			directory, err = os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		}
		listen := fmt.Sprintf("%s:%s", address, port)
		http.ListenAndServe(listen, http.FileServer(http.Dir(directory)))
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	serveCmd.PersistentFlags().StringVar(&directory, "dir", "", "The webroot to serve files from")
	serveCmd.PersistentFlags().StringVar(&port, "port", "8080", "The listen port")

	serveCmd.PersistentFlags().StringVar(&address, "addr", "0.0.0.0", "The listen address.")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
