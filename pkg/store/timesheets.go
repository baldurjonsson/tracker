package store

import (
	"encoding/json"
	"time"
)

type Timesheet struct {
	ID      int
	From    time.Time
	To      time.Time
	Name    string
	Profile *Profile
	Entries []*Entry
}

type TimesheetStore struct {
	filename   string
	Timesheets []*Timesheet
}

func NewTimesheetStore(filename string) *TimesheetStore {
	return &TimesheetStore{
		filename,
		[]*Timesheet{},
	}
}

func (t *TimesheetStore) GetFilename() string {
	return t.filename
}

func (t *TimesheetStore) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &t.Timesheets)
}

func (t *TimesheetStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Timesheets)
}
