package db

import (
	"ca-zoooom/entity"
	"ca-zoooom/util"
	"log"
	"time"
)

type TemplateRepository struct {
	SqlHandler
}

func (repo *TemplateRepository) Count() (count int, err error) {
	c, err := repo.SelectInt("select count(*) from templates")
	count = int(c)
	return
}

func (repo *TemplateRepository) Get(limit int, offset int, keyword string) (templates entity.Templates, err error) {
	//_, err = repo.Select(&templates,
	//	"select * from templates LEFT JOIN template_tags ON templates.id = template_tags.template_id LEFT JOIN tags ON template_tags.tag_id = tags.id where tags.title = ? limit ? offset ?",
	//	keyword, limit, offset)
	keywordLike := "%"+keyword+"%"
	_, err = repo.Select(&templates, "select templates.id, templates.uid, templates.background_url, templates.generated_sample_url, templates.created_at, templates.updated_at from templates " +
		"LEFT JOIN template_tags ON templates.id = template_tags.template_id " +
		"LEFT JOIN tags ON template_tags.tag_id = tags.id " +
		"where tags.title like ? " +
		"limit ? offset ?",
		keywordLike, limit, offset)
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

func (repo *TemplateRepository) GetByUniqueId(uid string) (template entity.Template, err error) {
	err = repo.SelectOne(&template, "select * from templates where uid=?", uid)
	if err != nil {
		// Logging
		return
	}
	return
}

func (repo *TemplateRepository) Insert(template *entity.Template) (err error) {
	template.Uid = util.GenerateUfid()
	template.UpdatedAt = time.Now()
	template.CreatedAt = time.Now()
	err = repo.SqlHandler.Insert(template)
	if err != nil {
		// Logging
		return
	}
	return
}
