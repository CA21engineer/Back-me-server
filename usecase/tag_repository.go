package usecase

import "ca-zoooom/entity"

type TagRepository interface {
	Get() (entity.Tags, error)
	GetById(int) (entity.Tag, error)
	GetByTitle(string) (entity.Tag, error)
	Insert(*entity.Tag) error
}
