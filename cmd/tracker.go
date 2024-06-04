package main

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/baldurjonsson/tracker/pkg/action"
	"github.com/baldurjonsson/tracker/pkg/store"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "tracker",
		Usage: "Track your time",
		Before: func(c *cli.Context) error {
			s := store.NewStore(get_store_dir())
			c.Context = context.WithValue(c.Context, "store", s)
			return s.Load()
		},
		Commands: []*cli.Command{
			{
				Name:    "profile",
				Usage:   "Manage profile",
				Aliases: []string{"p"},
				Action:  action.ShowProfile,
				Subcommands: []*cli.Command{
					{
						Name:    "show",
						Aliases: []string{"s"},
						Usage:   "Show Profile",
						Action:  action.ShowProfile,
					},
					{
						Name:    "set",
						Aliases: []string{"c"},
						Usage:   "Set Profile",
						Action:  action.SetProfile,
					},
				},
			},
			{
				Name:    "timesheets",
				Usage:   "Manage timesheets",
				Aliases: []string{"t"},
				Action:  action.ListTimesheets,
				Subcommands: []*cli.Command{
					{
						Name:    "list",
						Aliases: []string{"l"},
						Usage:   "List Timesheets",
						Action:  action.ListTimesheets,
					},
					{
						Name:    "read",
						Aliases: []string{"r"},
						Usage:   "Read Timesheet",
						Args:    true,
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:  "json",
								Usage: "Output as JSON",
							},
						},
						Action: action.ReadTimesheet,
					},
					{
						Name:    "create",
						Aliases: []string{"c"},
						Usage:   "Create Timesheet",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Usage:    "Name of the timesheet",
								Required: false,
							},
							&cli.TimestampFlag{
								Name:     "from",
								Usage:    "Start date of the timesheet",
								Required: true,
								Layout:   "2006-01-02",
							},
							&cli.TimestampFlag{
								Name:     "to",
								Usage:    "End date of the timesheet",
								Required: true,
								Layout:   "2006-01-02",
							},
							&cli.BoolFlag{
								Name:     "nonbillable",
								Usage:    "Include Non-Billable entries",
								Required: false,
							},
							&cli.StringFlag{
								Name:     "project",
								Usage:    "Project to filter entries by",
								Required: false,
							},
						},
						Action: action.CreateTimesheet,
					},
				},
			},
			{
				Name:    "entries",
				Usage:   "Manage time entries",
				Aliases: []string{"e"},
				Action:  action.ListEntries,
				Subcommands: []*cli.Command{
					{
						Name:    "list",
						Aliases: []string{"l"},
						Usage:   "List Entries",
						Action:  action.ListEntries,
					},
					{
						Name:    "create",
						Aliases: []string{"c"},
						Usage:   "Create Entry",
						Action:  action.CreateEntry,
					},
					{
						Name:    "delete",
						Aliases: []string{"d"},
						Usage:   "Delete Entry",
						Args:    true,
						Action:  action.DeleteEntry,
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func get_store_dir() string {
	STORE_DIR := os.Getenv("STORE_DIR")
	if STORE_DIR == "" {
		STORE_DIR = path.Join(os.Getenv("HOME"), ".tracker")
		if _, err := os.Stat(STORE_DIR); os.IsNotExist(err) {
			os.Mkdir(STORE_DIR, 0755)
		}
	}
	return STORE_DIR
}
