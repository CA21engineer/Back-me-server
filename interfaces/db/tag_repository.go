package db

import (
	"ca-zoooom/entity"
	"fmt"
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
	_, err = repo.Select(tag, "select * from tags where id=?", id)
	fmt.Println(tag)
	if err != nil {
		return
	}
	return
}

func (repo *TagRepository) IgnoreInsert(t *entity.Tag) (err error) {
	t.UpdatedAt = time.Now()
	t.CreatedAt = time.Now()
	// IgnoreInsertの実装はとりあえずあとまわし
	err = repo.SqlHandler.Insert(t)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}
