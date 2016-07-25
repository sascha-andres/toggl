package account

import (
	"fmt"
	"github.com/jason0x43/go-toggl"
	// "time"
)

// Dump writes out account data
func Dump(settingToken string, printTimeEntry bool) error {
	session := toggl.OpenSession(settingToken)
	account, err := session.GetAccount()
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("  Timezone: %s", account.Data.Timezone))
	fmt.Println(fmt.Sprintf("  Workspaces: %d", len(account.Data.Workspaces)))
	fmt.Println(fmt.Sprintf("  Projects: %d", len(account.Data.Projects)))
	fmt.Println(fmt.Sprintf("  Tags: %d", len(account.Data.Tags)))

	timeEntries := len(account.Data.TimeEntries)
	if printTimeEntry && timeEntries > 0 {
		if nil == account.Data.TimeEntries[timeEntries-1].Stop {
			fmt.Println(fmt.Sprintf("Current time entry: %s - Running: %s", account.Data.TimeEntries[timeEntries-1].Start.Format("15:04:05"), account.Data.TimeEntries[timeEntries-1].Description))
		} else {
			fmt.Println(fmt.Sprintf("Last time entry: %s - %s: %s", account.Data.TimeEntries[timeEntries-1].Start.Format("15:04:05"), account.Data.TimeEntries[timeEntries-1].Stop.Format("15:04:05"), account.Data.TimeEntries[timeEntries-1].Description))
		}
	}

	return nil
}
