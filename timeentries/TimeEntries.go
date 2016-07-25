package timeentries

import (
	"fmt"
	"github.com/jason0x43/go-toggl"
)

// NewTimeEntry creates a new time entry
func NewTimeEntry(settingToken, settingDescription, settingProjectName string) error {
	session := toggl.OpenSession(settingToken)
	if len(settingProjectName) == 0 {
		_, err := session.StartTimeEntry(settingDescription)
		if err != nil {
			return err
		}
	} else {
		// find project
		account, err := session.GetAccount()
		if err != nil {
			return err
		}
		var index = -1
		for i, prj := range account.Data.Projects {
			if prj.Name == settingProjectName {
				index = i
				break
			}
		}
		if -1 == index {
			fmt.Println("Project not found. Use list-projects to view them")
		} else {
			_, err = session.StartTimeEntryForProject(settingDescription, account.Data.Projects[index].ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// StopCurrent stops a running time entry creates a new time entry
func StopCurrent(settingToken string) error {
	session := toggl.OpenSession(settingToken)
	account, err := session.GetAccount()
	if err != nil {
		return err
	}
	for _, te := range account.Data.TimeEntries {
		if nil == te.Stop {
			session.StopTimeEntry(te)
			break
		}
	}

	return nil
}
