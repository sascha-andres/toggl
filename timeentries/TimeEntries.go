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

package timeentries

import (
	"fmt"

	"github.com/sascha-andres/go-toggl"
	"github.com/sascha-andres/toggl/types"
)

// New creates a new time entry
func New(settings types.Settings) error {
	toggl.DisableLogging()

	session := toggl.OpenSession(settings.Token)
	if len(settings.ProjectName) == 0 {
		_, err := session.StartTimeEntry(settings.Description)
		if err != nil {
			return err
		}
	} else {
		// find project
		account, err := session.GetAccount()
		if err != nil {
			return err
		}
		var index int
		if index, err = getProjectIndex(account, settings); err != nil {
			return err
		}
		_, err = session.StartTimeEntryForProject(settings.Description, account.Data.Projects[index].ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// StopCurrent stops a running time entry creates a new time entry
func StopCurrent(settingToken string) error {
	toggl.DisableLogging()

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

// Update sets new values
func Update(settings types.Settings) error {
	toggl.DisableLogging()

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

func getProjectIndex(account toggl.Account, settings types.Settings) (int, error) {
	for i, prj := range account.Data.Projects {
		if prj.Name == settings.ProjectName {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Project not found: %s", settings.ProjectName)
}

func getCurrentTimeEntry(account toggl.Account) (*toggl.TimeEntry, error) {
	for _, te := range account.Data.TimeEntries {
		if nil == te.Stop {
			return &te, nil
		}
	}
	return nil, fmt.Errorf("No current time entry")
}
