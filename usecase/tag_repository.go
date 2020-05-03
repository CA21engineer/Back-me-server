package usecase

import "ca-zoooom/entity"

type TagRepository interface {
	Get() (entity.Tags, error)
}
