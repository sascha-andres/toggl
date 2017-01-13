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

	"github.com/sascha-andres/toggl/account"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Print some account information",
	Long: `The command prints out the following information:
Timezone
Workspace count
Project count
Tag count

If you add --time the last time entry will be printed, too.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkPFlags()
		if err := account.Dump(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(accountCmd)
	accountCmd.Flags().BoolP("time", "t", false, "Print your last timeentry")
	viper.BindPFlag("account.time", accountCmd.Flags().Lookup("time"))
}
