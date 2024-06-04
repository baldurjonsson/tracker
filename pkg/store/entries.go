package store

import (
	"encoding/json"
	"time"
)

type EntriesStore struct {
	filename string
	Entries  []*Entry
}

func NewEntriesStore(filename string) *EntriesStore {
	return &EntriesStore{
		filename,
		[]*Entry{},
	}
}

func (e *EntriesStore) GetFilename() string {
	return e.filename
}
func (e *EntriesStore) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e.Entries)
}
func (e *EntriesStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Entries)
}

type Entry struct {
	ID       int
	Project  string
	Hours    float64
	Date     time.Time
	Notes    string
	Billable bool
}
