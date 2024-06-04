package store

import (
	"os"
	"path"
)

type StoreObject interface {
	GetFilename() string
	UnmarshalJSON(data []byte) error
	MarshalJSON() ([]byte, error)
}

type Store struct {
	dir        string
	Ids        *IdStore
	Profile    *ProfileStore
	Entries    *EntriesStore
	Timesheets *TimesheetStore
}

func NewStore(dir string) *Store {
	s := &Store{
		dir,
		NewIdStore("ids.json"),
		NewProfileStore("profile.json"),
		NewEntriesStore("entries.json"),
		NewTimesheetStore("timesheets.json"),
	}
	return s
}
func (s *Store) GetObjects() []StoreObject {
	return []StoreObject{s.Ids, s.Profile, s.Entries, s.Timesheets}
}

func (s *Store) Load() error {
	for _, so := range s.GetObjects() {
		data, err := os.ReadFile(path.Join(s.dir, so.GetFilename()))
		if err != nil { // If the file does not exist, we assume the store is empty
			continue
		}
		err = so.UnmarshalJSON(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) Save() error {
	for _, so := range s.GetObjects() {
		data, _ := so.MarshalJSON()
		filename := so.GetFilename()
		err := os.WriteFile(path.Join(s.dir, filename), data, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}
