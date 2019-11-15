/*
Copyright © 2019 PAVAN GARLAPATI <garlapatipavan94@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "reg",
	Short: "reg is mimic of reg.exe with limited operations",
	Long: `
		reg is mimic of Windows reg.exe CLI
		This tool is limited to QUERY, ADD, DELETE operations

		USAGE: 	reg [OPERATION] [options]
		Examples: 
				reg ADD [key] [value]
				reg QUERY [key]
				reg DELETE [key]
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println("lets see")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Adding Sub Commands (Query, Add, Delete)
	rootCmd.AddCommand(QueryCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
