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

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new time entry in toggl",
	Long: `To start a new time entry in toggl use toggl time start

You can provide a project using --project and a descriptionusing --desc.

Example toggl time start --desc "Hello toggl!"`,
	Run: func(cmd *cobra.Command, args []string) {
		checkPFlags()
		if "" == viper.GetString("time.start.description") {
			log.Fatal("A description is required")
		}
		if err := timeentries.New(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	timeCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("project", "p", "", "Assign project")
	startCmd.Flags().StringP("desc", "d", "", "Description")
	viper.BindPFlag("time.start.project", startCmd.Flags().Lookup("project"))
	viper.BindPFlag("time.start.description", startCmd.Flags().Lookup("desc"))
}
