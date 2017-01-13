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
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop working on running time entry",
	Long: `If you stop working on a timeentry you can stop the
time entry by calling

  toggl time stop`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := timeentries.StopCurrent(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	timeCmd.AddCommand(stopCmd)
}
