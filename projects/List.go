package projects

import (
	"fmt"

	toggl "github.com/jason0x43/go-toggl"
	"github.com/spf13/viper"
)

// List writes out project data
func List() error {
	session := toggl.OpenSession(viper.GetString("token"))
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
