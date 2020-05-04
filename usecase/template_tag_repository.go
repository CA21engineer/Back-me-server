package usecase

import "ca-zoooom/entity"

type TemplateTagRepository interface {
	GetByTemplateId(int) (entity.TemplateTags, error)
	GetByTagId(int) (entity.TemplateTags, error)
}
