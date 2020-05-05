package db

import (
	"ca-zoooom/entity"
	"log"
)

type TemplateTagRepository struct {
	SqlHandler
}

func (repo *TemplateTagRepository) GetByTemplateId(t int) (templateTags entity.TemplateTags, err error) {
	_, err = repo.Select(&templateTags, "select * from template_tags where template_id=?", t)
	if err != nil {
		return
	}
	return
}

func (repo *TemplateTagRepository) GetByTagId(t int) (templateTag entity.TemplateTag, err error) {
	err = repo.SelectOne(&templateTag, "select * from template_tags where tag_id=?", t)
	if err != nil {
		return
	}
	return
}

func (repo *TemplateTagRepository) Insert(t *entity.TemplateTag) (err error) {
	err = repo.SqlHandler.Insert(t)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}
