package timeentries

import (
	toggl "github.com/jason0x43/go-toggl"
	"github.com/spf13/viper"
)

// StopCurrent stops a running time entry creates a new time entry
func StopCurrent() error {
	session := toggl.OpenSession(viper.GetString("token"))
	account, err := session.GetAccount()
	if err != nil {
		return err
	}
	var timeEntry *toggl.TimeEntry
	timeEntry, _ = getCurrentTimeEntry(account)
	if nil != timeEntry {
		session.StopTimeEntry(*timeEntry)
	}

	return nil
}
