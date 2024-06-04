package action

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/baldurjonsson/tracker/pkg/store"
	"github.com/buger/goterm"
	"github.com/urfave/cli/v2"
)

func ListTimesheets(c *cli.Context) error {
	s := c.Context.Value("store").(*store.Store)
	for _, ts := range s.Timesheets.Timesheets {
		fmt.Println(ts)
	}
	return nil
}

func CreateTimesheet(c *cli.Context) error {
	s := c.Context.Value("store").(*store.Store)
	nonbillable := c.Bool("nonbillable")
	project := c.String("project")
	ts := &store.Timesheet{}
	s.Ids.Ids["timesheet"]++
	ts.ID = s.Ids.Ids["timesheet"]
	ts.From = *c.Timestamp("from")
	ts.To = c.Timestamp("to").Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	ts.Name = c.String("name")
	ts.Profile = s.Profile.Profile
	ts.Entries = make([]*store.Entry, 0)
	for _, e := range s.Entries.Entries {
		if e.Date.After(ts.From) && e.Date.Before(ts.To) {
			if (nonbillable || e.Billable) && (project == "" || e.Project == project) {
				ts.Entries = append(ts.Entries, e)
			}
		}
	}
	s.Timesheets.Timesheets = append(s.Timesheets.Timesheets, ts)
	data, _ := json.MarshalIndent(ts, "", "  ")
	fmt.Println(string(data))
	return s.Save()
}

func ReadTimesheet(c *cli.Context) error {
	s := c.Context.Value("store").(*store.Store)
	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return err
	}
	var ts *store.Timesheet
	for _, t := range s.Timesheets.Timesheets {
		if t.ID == id {
			ts = t
			break
		}
	}
	if ts == nil {
		return fmt.Errorf("timesheet with id %d not found", id)
	}
	if c.Bool("json") {
		data, _ := json.MarshalIndent(ts, "", "  ")
		fmt.Println(string(data))
		return nil
	}
	w := goterm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(w, "ID:\t%d\n", ts.ID)
	if ts.Name != "" {
		fmt.Fprintf(w, "Name:\t%s\n", ts.Name)
	}
	fmt.Fprintf(w, "From:\t%s\n", ts.From.Format("2006-01-02"))
	fmt.Fprintf(w, "To:\t%s\n", ts.To.Format("2006-01-02"))
	fmt.Fprintf(w, "Profile:\t%s<%s>\n", ts.Profile.Name, ts.Profile.Email)
	fmt.Fprintf(w, "Entries:\t%d\n", len(ts.Entries))
	goterm.Print(w)
	goterm.Flush()
	PrintEntriesTable(ts.Entries)

	return nil
}
