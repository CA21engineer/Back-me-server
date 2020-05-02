package entity

import "time"

type Videos []Video

type Video struct {
	Id           int       `json:"id" db:"id" `
	AdminId      int       `json:"admin_id" db:"admin_id"`
	Title        string    `json:"title" db:"title"`
	Note         string    `json:"note" db:"note"`
	Uri          string    `json:"uri" db:"uri"`
	ThumbnailUri string    `json:"thumbnail_uri" db:"thumbnail_uri"`
	Duration     int       `json:"duration" db:"duration"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}
