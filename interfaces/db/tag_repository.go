package db

import (
	"ca-zoooom/entity"
)

type TagRepository struct {
	SqlHandler
}

func (repo *TagRepository) Get() (tags entity.Tags, err error) {
	_, err = repo.Select(&tags, "select * from tags order by id")
	if err != nil {
		return
	}
	return
}
