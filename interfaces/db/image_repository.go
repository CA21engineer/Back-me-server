package db

import (
	"ca-zoooom/entity"
)

type ImageRepository struct {
	SqlHandler
}

func (repo *ImageRepository) Get() (images entity.Images, err error) {
	// タグとかフィルタする系のクエリパラメータ
	_, err = repo.Select(&images, "select * from images order by id")
	if err != nil {
		return
	}
	return
}

func (repo *ImageRepository) GetById(id int) (image entity.Image, err error) {
	err = repo.SelectOne(&image, "select * from images where id=?", id)
	if err != nil {
		return
	}
	return
}

func (repo *ImageRepository) Insert(v *entity.Image) (err error) {
	err = repo.SqlHandler.Insert(v)
	if err != nil {
		return
	}
	return
}
