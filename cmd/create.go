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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Long: `Use this command to create a new project.

A project in toggl is a way to group time entries.

Example: toggl project create --name "My project"

Note: --name is required`,
	Run: func(cmd *cobra.Command, args []string) {
		if "" == viper.GetString("project.name") {
			log.Fatal("Please provide --name")
		}
		if err := projects.Add(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	projectCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Project name to add")
	viper.BindPFlag("project.name", createCmd.Flags().Lookup("name"))
}
