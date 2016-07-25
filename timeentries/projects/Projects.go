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
