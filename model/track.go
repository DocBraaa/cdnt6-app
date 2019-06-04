package model

// Track represents a track of album.
type Track struct {
	ID     string `json:"id,omitempty" bson:"_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Status bool   `json:"status"`
}
