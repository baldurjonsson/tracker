package action

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/baldurjonsson/tracker/pkg/store"
	"github.com/buger/goterm"
	"github.com/urfave/cli/v2"
)

func PrintEntriesTable(entries []*store.Entry) {
	entriesTable := goterm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprint(entriesTable, "ID\tProject\tDate\tHours\tBillable\tNotes\n")
	for _, entry := range entries {
		billableChar := '✓'
		if entry.Billable == false {
			billableChar = '✗'
		}
		fmt.Fprintf(entriesTable, "%d\t%s\t%s\t%.1f\t%c\t%s\n", entry.ID, entry.Project, entry.Date.Format("2006-01-02"), entry.Hours, billableChar, entry.Notes)
	}
	goterm.Println(entriesTable)
	goterm.Flush()
}

func ListEntries(c *cli.Context) error {
	s := c.Context.Value("store").(*store.Store)
	from := c.Timestamp("from")

	if from != nil {
		var filtered []*store.Entry
		for _, entry := range s.Entries.Entries {
			if entry.Date.After(*from) {
				filtered = append(filtered, entry)
			}
		}
		PrintEntriesTable(filtered)
	} else {
		PrintEntriesTable(s.Entries.Entries)
	}
	return nil
}

func CreateEntry(c *cli.Context) error {
	b := bufio.NewReader(os.Stdin)
	s := c.Context.Value("store").(*store.Store)
	project := prompt(b, "Project")
	notes := prompt(b, "Notes")
	hours, _ := strconv.ParseFloat(prompt(b, "Hours"), 64)
	dateStr := prompt(b, "Date")
	var date time.Time
	if dateStr == "" {
		date = time.Now()
	} else {
		date, _ = time.Parse("2006-01-02", dateStr)
	}
	billableStr := prompt(b, "Billable")
	billable := true
	if billableStr != "" && billableStr[:1] == "n" {
		billable = false
	}
	s.Ids.Ids["entry"]++
	entry := &store.Entry{
		ID:       s.Ids.Ids["entry"],
		Project:  project,
		Hours:    hours,
		Date:     date,
		Notes:    notes,
		Billable: billable,
	}
	s.Entries.Entries = append(s.Entries.Entries, entry)
	SortEntries(c)
	return s.Save()
}

func SortEntries(c *cli.Context) error {
	s := c.Context.Value("store").(*store.Store)
	sort.Slice(s.Entries.Entries, func(i, j int) bool {
		return s.Entries.Entries[i].Date.Before(s.Entries.Entries[j].Date)
	})
	return s.Save()
}

func DeleteEntry(c *cli.Context) error {
	s := c.Context.Value("store").(*store.Store)
	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return err
	}
	for i, entry := range s.Entries.Entries {
		if entry.ID == id {
			s.Entries.Entries = append(s.Entries.Entries[:i], s.Entries.Entries[i+1:]...)
			break
		}
	}
	return s.Save()
}
