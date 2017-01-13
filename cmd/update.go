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

	"github.com/sascha-andres/toggl/timeentries"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a running time entry",
	Long: `Amend a running time entry with a description of a project or change the values.

  toggl time update --desc "New description" --project "My project"`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := timeentries.Update(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	timeCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("project", "p", "", "Assign project")
	updateCmd.Flags().StringP("desc", "d", "", "Description")
	viper.BindPFlag("time.project", updateCmd.Flags().Lookup("project"))
	viper.BindPFlag("time.description", updateCmd.Flags().Lookup("desc"))
}
