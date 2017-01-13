package projects

import (
	toggl "github.com/jason0x43/go-toggl"
	"github.com/sascha-andres/toggl/types"
)

// Add a new project
func Add(settings types.Settings) error {
	session := toggl.OpenSession(settings.Token)
	account, err := session.GetAccount()
	if err != nil {
		return err
	}
	wid := account.Data.Workspaces[0].ID
	_, err = session.CreateProject(settings.ProjectName, wid)
	return err
}
