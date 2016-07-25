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
