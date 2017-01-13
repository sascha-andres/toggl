package timeentries

import (
	toggl "github.com/jason0x43/go-toggl"
	"github.com/sascha-andres/toggl/types"
)

// Update sets new values
func Update(settings types.Settings) error {
	session := toggl.OpenSession(settings.Token)

	account, err := session.GetAccount()
	if err != nil {
		return err
	}

	var timeEntry *toggl.TimeEntry
	if timeEntry, err = getCurrentTimeEntry(account); err != nil {
		return err
	}

	timeEntry.Description = settings.Description

	if 0 < len(settings.ProjectName) {
		var index int
		if index, err = getProjectIndex(account, settings); err != nil {
			return err
		}
		timeEntry.Pid = account.Data.Projects[index].ID
	}

	_, err = session.UpdateTimeEntry(*timeEntry)

	return err
}
