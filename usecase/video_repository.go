package usecase

import "ca-zoooom/entity"

type VideoRepository interface {
	Get() (entity.Videos, error)
	GetById(int) (entity.Video, error)
	Insert(*entity.Video) error
	Update(*entity.Video) (int64, error)
	Delete(*entity.Video) error
}
