package usecase

import "ca-zoooom/entity"

type ImageRepository interface {
	Get() (entity.Images, error)
	GetById(int) (entity.Image, error)
	Insert(*entity.Image) error
	Update(*entity.Image) (int64, error)
	Delete(*entity.Image) error
}
