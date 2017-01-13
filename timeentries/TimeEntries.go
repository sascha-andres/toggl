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

func getProjectIndex(account toggl.Account, project string) (int, error) {
	for i, prj := range account.Data.Projects {
		if prj.Name == project {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Project not found: %s", project)
}

func getCurrentTimeEntry(account toggl.Account) (*toggl.TimeEntry, error) {
	for _, te := range account.Data.TimeEntries {
		if nil == te.Stop {
			return &te, nil
		}
	}
	return nil, fmt.Errorf("No current time entry")
}
