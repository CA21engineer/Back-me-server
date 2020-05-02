package usecase

import (
	"ca-zoooom/entity"
)

type VideoInteractor struct {
	VideoRepository VideoRepository
	StatusCode      int
}

func (interactor *VideoInteractor) ListVideos() (v entity.Videos, err error) {
	v, err = interactor.VideoRepository.Get()
	interactor.StatusCode = 200
	return
}

func (interactor *VideoInteractor) GetByID(id int) (v entity.Video, err error) {
	v, err = interactor.VideoRepository.GetById(id)
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	interactor.StatusCode = 200
	return
}

func (interactor *VideoInteractor) Add(v *entity.Video) (video entity.Video, err error) {
	err = interactor.VideoRepository.Insert(v)
	if err != nil {
		return
	}
	video, err = interactor.VideoRepository.GetById(v.Id)
	interactor.StatusCode = 201
	return
}

func (interactor *VideoInteractor) Update(new *entity.Video, id int) (video entity.Video, err error) {
	v, err := interactor.VideoRepository.GetById(id)
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	// Update params ToDo: When new param is nil...
	v.Title = new.Title
	v.Note = new.Note
	v.Uri = new.Uri
	v.ThumbnailUri = new.ThumbnailUri
	v.Duration = new.Duration
	// DB
	_, err = interactor.VideoRepository.Update(&v)
	if err != nil {
		interactor.StatusCode = 500
		return
	}
	video, err = interactor.VideoRepository.GetById(id)
	interactor.StatusCode = 200
	return
}

func (interactor *VideoInteractor) Delete(id int) (err error) {
	video, err := interactor.VideoRepository.GetById(id)
	err = interactor.VideoRepository.Delete(&video)
	if err != nil {
		return
	}
	interactor.StatusCode = 204
	return
}
