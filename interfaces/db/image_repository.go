package db

import (
	"ca-zoooom/entity"
	"time"
)

type ImageRepository struct {
	SqlHandler
}

func (repo *ImageRepository) Count() (count int, err error) {
	c, err := repo.SelectInt("select count(*) from images where is_private=false")
	count = int(c)
	return
}

func (repo *ImageRepository) Get(limit int, offset int) (images entity.Images, err error) {
	_, err = repo.Select(&images, "select * from images where is_private=false order by id desc limit ? offset ?", limit, offset)
	if err != nil {
		// Logging
		return
	}
	return
}

func (repo *ImageRepository) GetById(id int) (image entity.Image, err error) {
	err = repo.SelectOne(&image, "select * from images where id=?", id)
	if err != nil {
		// Logging
		return
	}
	return
}

func (repo *ImageRepository) Insert(image *entity.Image) (err error) {
	image.UpdatedAt = time.Now()
	image.CreatedAt = time.Now()
	err = repo.SqlHandler.Insert(image)
	if err != nil {
		// Logging
		return
	}
	return
}
