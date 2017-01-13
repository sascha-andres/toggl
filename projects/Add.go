package projects

import (
	toggl "github.com/jason0x43/go-toggl"
	"github.com/spf13/viper"
)

// Add a new project
func Add() error {
	session := toggl.OpenSession(viper.GetString("token"))
	account, err := session.GetAccount()
	if err != nil {
		return err
	}
	wid := account.Data.Workspaces[0].ID
	_, err = session.CreateProject(viper.GetString("project.name"), wid)
	return err
}
