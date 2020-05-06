package usecase

import "ca-zoooom/entity"

type ImageRepository interface {
	Count() (int, error)
	Get(limit int, offset int) (entity.Images, error)
	GetById(int) (entity.Image, error)
	Insert(*entity.Image) error
}
