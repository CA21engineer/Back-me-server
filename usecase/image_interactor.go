package usecase

import (
	"ca-zoooom/entity"
)

type ImageInteractor struct {
	ImageRepository ImageRepository
	StatusCode      int
}

func (interactor *ImageInteractor) ListImages() (v entity.Images, err error) {
	v, err = interactor.ImageRepository.Get()
	interactor.StatusCode = 200
	return
}

func (interactor *ImageInteractor) GetByID(id int) (v entity.Image, err error) {
	v, err = interactor.ImageRepository.GetById(id)
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	interactor.StatusCode = 200
	return
}

func (interactor *ImageInteractor) Add(v *entity.Image) (image entity.Image, err error) {
	err = interactor.ImageRepository.Insert(v)
	if err != nil {
		return
	}
	image, err = interactor.ImageRepository.GetById(v.Id)
	interactor.StatusCode = 201
	return
}
