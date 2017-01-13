package timeentries

import (
	toggl "github.com/jason0x43/go-toggl"
	"github.com/sascha-andres/toggl/types"
)

// New creates a new time entry
func New(settings types.Settings) error {
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
