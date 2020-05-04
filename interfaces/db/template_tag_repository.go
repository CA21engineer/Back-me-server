package db

import (
	"ca-zoooom/entity"
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

func (repo *TemplateTagRepository) GetByTagId(t int) (templateTags entity.TemplateTags, err error) {
	_, err = repo.Select(&templateTags, "select * from template_tags where tag_id=?", t)
	if err != nil {
		return
	}
	return
}
