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

	"livingit.de/code/toggl/projects"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your projects",
	Long: `Show a list of your projects. Displays one project per line.
	
Format used: "  <Name> (<ID>)"

toggl project list`,
	Run: func(cmd *cobra.Command, args []string) {
		checkPFlags()
		if err := projects.List(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	projectCmd.AddCommand(listCmd)
}
