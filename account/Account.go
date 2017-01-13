// Copyright 2016 Sascha Andres

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package account

import (
	"fmt"

	"github.com/jason0x43/go-toggl"
	"github.com/spf13/viper"
)

// Dump writes out account data
func Dump() error {
	session := toggl.OpenSession(viper.GetString("token"))
	account, err := session.GetAccount()
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("  Timezone: %s", account.Data.Timezone))
	fmt.Println(fmt.Sprintf("  Workspaces: %d", len(account.Data.Workspaces)))
	fmt.Println(fmt.Sprintf("  Projects: %d", len(account.Data.Projects)))
	fmt.Println(fmt.Sprintf("  Tags: %d", len(account.Data.Tags)))

	timeEntries := len(account.Data.TimeEntries)
	if viper.GetBool("account.time") && timeEntries > 0 {
		if nil == account.Data.TimeEntries[timeEntries-1].Stop {
			fmt.Println(fmt.Sprintf("Current time entry: %s - Running: %s", account.Data.TimeEntries[timeEntries-1].Start.Format("15:04:05"), account.Data.TimeEntries[timeEntries-1].Description))
		} else {
			fmt.Println(fmt.Sprintf("Last time entry: %s - %s: %s", account.Data.TimeEntries[timeEntries-1].Start.Format("15:04:05"), account.Data.TimeEntries[timeEntries-1].Stop.Format("15:04:05"), account.Data.TimeEntries[timeEntries-1].Description))
		}
	}

	return nil
}
