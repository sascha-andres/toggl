package timeentries

import (
	toggl "github.com/jason0x43/go-toggl"
	"github.com/spf13/viper"
)

// Update sets new values
func Update() error {
	session := toggl.OpenSession(viper.GetString("token"))

	account, err := session.GetAccount()
	if err != nil {
		return err
	}

	var timeEntry *toggl.TimeEntry
	if timeEntry, err = getCurrentTimeEntry(account); err != nil {
		return err
	}

	timeEntry.Description = viper.GetString("time.update.description")

	if 0 < len(viper.GetString("time.update.project")) {
		var index int
		if index, err = getProjectIndex(account, viper.GetString("time.update.project")); err != nil {
			return err
		}
		timeEntry.Pid = account.Data.Projects[index].ID
	}

	_, err = session.UpdateTimeEntry(*timeEntry)

	return err
}
