package timeentries

import (
	toggl "github.com/jason0x43/go-toggl"
	"github.com/spf13/viper"
)

// New creates a new time entry
func New() error {
	session := toggl.OpenSession(viper.GetString("token"))
	if viper.GetString("time.start.project") == "" {
		_, err := session.StartTimeEntry(viper.GetString("time.start.description"))
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
		if index, err = getProjectIndex(account, viper.GetString("time.start.project")); err != nil {
			return err
		}
		_, err = session.StartTimeEntryForProject(viper.GetString("time.start.description"), account.Data.Projects[index].ID)
		if err != nil {
			return err
		}
	}

	return nil
}
