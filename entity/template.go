package entity

import "time"

type Templates []Template

type Template struct {
	Id                 int       `json:"id" db:"id" `
	Uid                string    `json:"uid" db:"uid"`
	DesignPatternId    int       `json:"design_pattern_id" db:"design_pattern_id"`
	IsPrivate          bool      `json:"is_private" db:"is_private"`
	BackGroundUrl      string    `json:"background_url" db:"background_url"`
	GeneratedSampleUrl string    `json:"generated_sample_url" db:"generated_sample_url"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
}
