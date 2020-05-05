package usecase

import "ca-zoooom/entity"

type TemplateRepository interface {
	Count() (int, error)
	Get(limit int, offset int, keyword string) (entity.Templates, error)
	GetById(int) (entity.Template, error)
	GetByUniqueId(string) (entity.Template, error)
	Insert(*entity.Template) error
}
