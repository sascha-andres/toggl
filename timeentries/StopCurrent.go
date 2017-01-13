package timeentries

import toggl "github.com/jason0x43/go-toggl"

// StopCurrent stops a running time entry creates a new time entry
func StopCurrent(settingToken string) error {
	session := toggl.OpenSession(settingToken)
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
