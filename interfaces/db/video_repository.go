package db

import (
	"ca-zoooom/entity"
)

type VideoRepository struct {
	SqlHandler
}

func (repo *VideoRepository) Get() (videos entity.Videos, err error) {
	_, err = repo.Select(&videos, "select * from videos order by id")
	if err != nil {
		return
	}
	return
}

func (repo *VideoRepository) GetById(id int) (video entity.Video, err error) {
	err = repo.SelectOne(&video, "select * from videos where id=?", id)
	if err != nil {
		return
	}
	return
}

func (repo *VideoRepository) Insert(v *entity.Video) (err error) {
	err = repo.SqlHandler.Insert(v)
	if err != nil {
		return
	}
	return
}

func (repo *VideoRepository) Update(v *entity.Video) (id int64, err error) {
	id, err = repo.SqlHandler.Update(v)
	if err != nil {
		return
	}

	return
}

func (repo *VideoRepository) Delete(v *entity.Video) error {
	_, err := repo.SqlHandler.Delete(v)
	if err != nil {
		return err
	}
	return err
}
