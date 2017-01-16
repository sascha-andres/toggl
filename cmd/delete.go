// Copyright © 2017 Sascha Andres <sascha.andres@outlook.com>
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
	"log"

	"github.com/sascha-andres/toggl/projects"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project for toggl",
	Long: `Delete a project on toggl.comp

Usage: toggl project delete --name "My project"

Note: --name is required`,
	Run: func(cmd *cobra.Command, args []string) {
		checkPFlags()
		if "" == viper.GetString("project.delete.name") {
			log.Fatal("Please provide --name")
		}
		if err := projects.Delete(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	projectCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("name", "n", "", "Project name to add")
	viper.BindPFlag("project.delete.name", deleteCmd.Flags().Lookup("name"))
}
