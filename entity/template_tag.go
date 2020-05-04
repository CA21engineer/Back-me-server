package entity

type TemplateTags []TemplateTag

type TemplateTag struct {
	Id         int `json:"id" db:"id"`
	TagId      int `json:"tag_id" db:"tag_id"`
	TemplateId int `json:"template_id" db:"template_id"`
}
