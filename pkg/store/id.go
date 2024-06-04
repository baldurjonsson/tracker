package store

import "encoding/json"

type IdStore struct {
	filename string
	Ids      map[string]int
}

func NewIdStore(filename string) *IdStore {
	return &IdStore{
		filename,
		map[string]int{
			"entry":     0,
			"timesheet": 0,
		},
	}
}
func (i *IdStore) GetFilename() string {
	return i.filename
}
func (i *IdStore) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &i.Ids)
}
func (i *IdStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Ids)
}
