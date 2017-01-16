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
	"github.com/jason0x43/go-toggl"
	"github.com/spf13/viper"
)

// Delete a  project
func Delete() error {
	session := toggl.OpenSession(viper.GetString("token"))
	account, err := session.GetAccount()
	if err != nil {
		return err
	}
	var projectToDelete toggl.Project
	for _, prj := range account.Data.Projects {
		if prj.Name == viper.GetString("project.delete.name") {
			projectToDelete = prj
			break
		}
	}
	_, err = session.DeleteProject(projectToDelete)
	return err
}
