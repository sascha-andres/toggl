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
	"log"
	"os"

	toggl "github.com/jason0x43/go-toggl"
	"github.com/sascha-andres/toggl/account"
	"github.com/sascha-andres/toggl/projects"
	"github.com/sascha-andres/toggl/timeentries"
	"github.com/sascha-andres/toggl/types"
	"github.com/urfave/cli"
)

var settings types.Settings

func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "token",
			Usage:       "Provide your API token",
			EnvVar:      "TOGGL_TOKEN",
			Destination: &settings.Token,
		},
	}

	commands := getAccountCommands()
	for _, cmd := range getProjectCommands() {
		commands = append(commands, cmd)
	}
	for _, cmd := range getTimeCommands() {
		commands = append(commands, cmd)
	}

	app.Commands = commands

	app.Name = "toggl"
	app.Version = "20170111"
	app.Usage = "A commandline toggl client"

	app.Run(os.Args)
}

func getAccountCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "account",
			Usage: "Dump account info",
			Action: func(c *cli.Context) error {
				return account.Dump(settings)
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "time",
					Usage:       "specify if you want to print your last timeentry",
					Destination: &settings.AccountLastTimeEntry,
				},
			},
		},
	}
}

func getProjectCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "project",
			Usage: "Work on projects",
			Subcommands: []cli.Command{
				{
					Name:  "list",
					Usage: "Lists all projects",
					Action: func(c *cli.Context) error {
						return projects.List(settings.Token)
					},
				},
				{
					Name:  "create",
					Usage: "Add a new project",
					Action: func(c *cli.Context) error {
						if 0 == len(settings.ProjectName) {
							log.Fatal("You have to provide a project (--name <project>)")
						}
						return projects.Add(settings)
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "name",
							Usage:       "project name",
							Destination: &settings.ProjectName,
						},
					},
				},
				{
					Name:  "delete",
					Usage: "Delete a project",
					Action: func(c *cli.Context) error {
						if 0 == len(settings.ProjectName) {
							log.Fatal("You have to provide a project (--name <project>)")
						}
						return projects.Delete(settings)
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "name",
							Usage:       "project name",
							Destination: &settings.ProjectName,
						},
					},
				},
			},
		},
	}
}

func getTimeCommands() []cli.Command {
	return []cli.Command{
		{
			Name:  "time",
			Usage: "Work on projects",
			Subcommands: []cli.Command{
				{
					Name:  "start",
					Usage: "Start time entry",
					Action: func(c *cli.Context) error {
						if len(settings.Description) == 0 {
							log.Fatal("You have to provide a description")
						}
						return timeentries.New(settings)
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "project",
							Usage:       "Assign project",
							Destination: &settings.ProjectName,
						},
						cli.StringFlag{
							Name:        "desc",
							Usage:       "Description",
							Destination: &settings.Description,
						},
					},
				},
				{
					Name:  "update",
					Usage: "Update a running time entry",
					Action: func(c *cli.Context) error {
						if len(settings.Description) == 0 {
							log.Fatal("You have to provide a description")
						}
						return timeentries.Update(settings)
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "project",
							Usage:       "Assign project",
							Destination: &settings.ProjectName,
						},
						cli.StringFlag{
							Name:        "desc",
							Usage:       "Description",
							Destination: &settings.Description,
						},
					},
				},
				{
					Name:  "stop",
					Usage: "Stop a running time entry",
					Action: func(c *cli.Context) error {
						return timeentries.StopCurrent(settings.Token)
					},
				},
			},
		},
	}
}

func init() {
	toggl.DisableLog()
}
