package store

import (
	"github.com/music/model"
)

// TrackStore is a CRUD.
type TrackStore interface {
	All() ([]model.Track, error)

	Create(*model.Track) error

	Find(string) (*model.Track, error)

	Update(*model.Track) error

	Delete(string) error
}
