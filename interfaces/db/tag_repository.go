package db

import (
	"ca-zoooom/entity"
	"log"
	"time"
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

func (repo *TagRepository) GetById(id int) (tag entity.Tag, err error) {
	err = repo.SelectOne(&tag, "select * from tags where id=?", id)
	if err != nil {
		return
	}
	return
}

func (repo *TagRepository) GetByTitle(title string) (tag entity.Tag, err error) {
	err = repo.SelectOne(&tag, "select * from tags where title=? limit 1", title)
	if err != nil {
		return
	}
	return
}

func (repo *TagRepository) Insert(t *entity.Tag) (err error) {
	t.UpdatedAt = time.Now()
	t.CreatedAt = time.Now()

	err = repo.SqlHandler.Insert(t)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}
