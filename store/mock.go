package store

import (
	"errors"
	"strconv"
	"time"

	"github.com/music/model"
)

type MockStore map[string]model.Track

// All returns all tracks.
func (s MockStore) All() ([]model.Track, error) {
	var data []model.Track

	for _, track := range s {
		data = append(data, track)
	}

	return data, nil
}

// Create save a new track.
func (s MockStore) Create(t *model.Track) error {
	timestamp := time.Now().UnixNano()
	id := strconv.FormatInt(timestamp, 10)

	t.ID = id
	s[id] = *t

	return nil
}

// Find finds a track by its ID.
func (s MockStore) Find(ID string) (*model.Track, error) {
	if track, ok := s[ID]; ok {
		return &track, nil
	}

	return nil, errors.New("track not found")
}

// Update updates a track.
func (s MockStore) Update(t *model.Track) error {
	return nil
}

// Delete deletes a track by its ID.
func (s MockStore) Delete(ID string) error {
	delete(s, ID)
	return nil
}
