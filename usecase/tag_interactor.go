package usecase

import (
	"ca-zoooom/entity"
)

type TagInteractor struct {
	TagRepository TagRepository
	StatusCode    int
}

func (interactor *TagInteractor) ListTags() (v entity.Tags, err error) {
	v, err = interactor.TagRepository.Get()
	interactor.StatusCode = 200
	return
}
