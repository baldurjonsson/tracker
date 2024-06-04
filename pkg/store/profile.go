package store

import "encoding/json"

type Profile struct {
	Name  string
	Email string
}

type ProfileStore struct {
	filename string
	Profile  *Profile
}

func NewProfileStore(filename string) *ProfileStore {
	return &ProfileStore{
		filename,
		&Profile{},
	}
}

func (p *ProfileStore) GetFilename() string {
	return p.filename
}

func (p *ProfileStore) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &p.Profile)
}

func (p *ProfileStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Profile)
}
