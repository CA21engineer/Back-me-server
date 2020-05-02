package entity

import (
	"time"
)

type Images []Image

type Image struct {
	Id        int       `json:"id" db:"id"`
	AssetUrl  string    `json:"asset_url" db:"asset_url"`
	IsPrivate bool      `json:"is_private" db:"is_private"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
