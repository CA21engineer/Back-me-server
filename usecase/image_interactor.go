package usecase

import (
	"ca-zoooom/entity"
)

type ImageInteractor struct {
	ImageRepository ImageRepository
	StatusCode      int
}

func (interactor *ImageInteractor) ListImages(limit int, offset int) (i entity.Images, totalPages int, err error) {
	i, err = interactor.ImageRepository.Get(limit, offset)
	totalPages, err = interactor.ImageRepository.Count()
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	interactor.StatusCode = 200
	return
}

func (interactor *ImageInteractor) GetByID(id int) (i entity.Image, err error) {
	i, err = interactor.ImageRepository.GetById(id)
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	interactor.StatusCode = 200
	return
}

func (interactor *ImageInteractor) Add(image *entity.Image) (i entity.Image, err error) {
	err = interactor.ImageRepository.Insert(image)
	if err != nil {
		interactor.StatusCode = 500
		return
	}
	i, err = interactor.ImageRepository.GetById(image.Id)
	if err != nil {
		interactor.StatusCode = 500
		return
	}
	interactor.StatusCode = 201
	return
}
