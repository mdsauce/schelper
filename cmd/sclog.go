// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/mdsauce/schelper/logger"
	"github.com/spf13/cobra"
)

// sclogCmd represents the sclog command
var sclogCmd = &cobra.Command{
	Use:   "sclog",
	Short: "Pass a valid file w/ the path to start the analysis.",
	Long: `The sclog command will read in the specified file line-by-line.
Use a full or relative path when specifying a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetupLogfile()
		if checkArgs(args) != true {
			logger.Disklog.Warn("Exiting")
			os.Exit(1)
		}
		fmt.Println("sclog called.  Looking for", args[0])
	},
}

func init() {
	rootCmd.AddCommand(sclogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sclogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sclogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkArgs(args []string) bool {
	if len(args) < 1 {
		logger.Disklog.Warn("Not enough arguments, no SC logfile specified")
		return false
	}
	return true
}
