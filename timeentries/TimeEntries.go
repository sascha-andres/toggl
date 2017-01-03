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
		var index = getProjectIndex(account, settings)
		if -1 == index {
			fmt.Println("Project not found. Use list-projects to view them")
		} else {
			_, err = session.StartTimeEntryForProject(settings.Description, account.Data.Projects[index].ID)
			if err != nil {
				return err
			}
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
	for _, te := range account.Data.TimeEntries {
		if nil == te.Stop {
			session.StopTimeEntry(te)
			break
		}
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

	var timeEntry toggl.TimeEntry

	for _, te := range account.Data.TimeEntries {
		if nil == te.Stop {
			timeEntry = te
			break
		}
	}

	timeEntry.Description = settings.Description

	if 0 < len(settings.ProjectName) {
		index := getProjectIndex(account, settings)
		if index == -1 {
			return fmt.Errorf("Project not found: %s", settings.ProjectName)
		}
		timeEntry.Pid = account.Data.Projects[index].ID
	}

	_, err = session.UpdateTimeEntry(timeEntry)

	return err
}

func getProjectIndex(account toggl.Account, settings types.Settings) int {
	for i, prj := range account.Data.Projects {
		if prj.Name == settings.ProjectName {
			return i
		}
	}
	return -1
}
