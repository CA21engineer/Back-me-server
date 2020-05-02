package db

import (
	"ca-zoooom/entity"
)

type ImageRepository struct {
	SqlHandler
}

func (repo *ImageRepository) Get() (images entity.Images, err error) {
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

func (repo *ImageRepository) Update(v *entity.Image) (id int64, err error) {
	id, err = repo.SqlHandler.Update(v)
	if err != nil {
		return
	}

	return
}

func (repo *ImageRepository) Delete(v *entity.Image) error {
	_, err := repo.SqlHandler.Delete(v)
	if err != nil {
		return err
	}
	return err
}
