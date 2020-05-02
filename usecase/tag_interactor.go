package usecase

import (
	"ca-zoooom/entity"
)

type tagInteractor struct {
	tagRepository tagRepository
	StatusCode      int
}

func (interactor *tagInteractor) ListTags() (v entity.Tags, err error) {
	v, err = interactor.TagRepository.Get()
	interactor.StatusCode = 200
	return
}
