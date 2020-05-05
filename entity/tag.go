package entity

import "time"

type Tags []Tag

type Tag struct {
	Id        int       `json:"id" db:"id" `
	Title     string    `json:"title" db:"title"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
