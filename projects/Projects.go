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

package projects

import (
	"fmt"
	"github.com/jason0x43/go-toggl"
)

// List writes out project data
func List(settingToken string) error {
	session := toggl.OpenSession(settingToken)
	account, err := session.GetAccount()
	if err != nil {
		return err
	}

	for _, prj := range account.Data.Projects {
		if prj.IsActive() {
			fmt.Println(fmt.Sprintf(" %s (%d)", prj.Name, prj.ID))
		}
	}

	return nil
}

// Add a new project
func Add(settingToken, settingProjectName string) error {
	session := toggl.OpenSession(settingToken)
	account, err := session.GetAccount()
	if err != nil {
		return err
	}
	wid := account.Data.Workspaces[0].ID
	_, err = session.CreateProject(settingProjectName, wid)
	return err
}

// Delete a  project
func Delete(settingToken, settingProjectName string) error {
	session := toggl.OpenSession(settingToken)
	account, err := session.GetAccount()
	if err != nil {
		return err
	}
	var projectToDelete toggl.Project
	for _, prj := range account.Data.Projects {
		if prj.Name == settingProjectName {
			projectToDelete = prj
			break
		}
	}
	_, err = session.DeleteProject(projectToDelete)
	return err
}
