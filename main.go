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

package main

import (
	"github.com/sascha-andres/toggl/account"
	"github.com/sascha-andres/toggl/projects"
	"github.com/sascha-andres/toggl/timeentries"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	settingVerbose, settingAccountLastTimeEntry          bool
	settingToken, settingProjectName, settingDescription string
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "token",
			Usage:       "Provide you API token",
			EnvVar:      "TOGGL_TOKEN",
			Destination: &settingToken,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "account",
			Usage: "Dump account info",
			Action: func(c *cli.Context) error {
				return account.Dump(settingToken, settingAccountLastTimeEntry)
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "time",
					Usage:       "specify if you want to print your last timeentry",
					Destination: &settingAccountLastTimeEntry,
				},
			},
		},
		{
			Name:  "project",
			Usage: "Work on projects",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "Lists all projects",
					Action: func(c *cli.Context) error {
						return projects.List(settingToken)
					},
				},
				{
					Name:  "create",
					Usage: "Add a new project",
					Action: func(c *cli.Context) error {
						return projects.Add(settingToken, settingProjectName)
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "name",
							Usage:       "project name",
							Destination: &settingProjectName,
						},
					},
				},
				{
					Name:  "delete",
					Usage: "Delete a project",
					Action: func(c *cli.Context) error {
						return projects.Delete(settingToken, settingProjectName)
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "name",
							Usage:       "project name",
							Destination: &settingProjectName,
						},
					},
				},
			},
		},
		{
			Name:  "time",
			Usage: "Work on projects",
			Subcommands: []cli.Command{
				{
					Name:  "start",
					Usage: "Start time entry",
					Action: func(c *cli.Context) error {
						if len(settingDescription) == 0 {
							log.Fatal("You have to provide a description")
						}
						return timeentries.NewTimeEntry(settingToken, settingDescription, settingProjectName)
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "project",
							Usage:       "Assign project",
							Destination: &settingProjectName,
						},
						cli.StringFlag{
							Name:        "desc",
							Usage:       "Description",
							Destination: &settingDescription,
						},
					},
				},
				{
					Name:  "stop",
					Usage: "Stop a running time entry",
					Action: func(c *cli.Context) error {
						return timeentries.StopCurrent(settingToken)
					},
				},
			},
		},
	}

	app.Name = "toggl"
	app.Version = "0.2"
	app.Usage = "A commandline toggl client"

	app.Run(os.Args)
}
