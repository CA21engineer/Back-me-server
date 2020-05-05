package usecase

import (
	"ca-zoooom/entity"
	"fmt"
	"log"
)

type TemplateInteractor struct {
	TemplateRepository TemplateRepository
	TagRepository 	   TagRepository
	TemplateTagRepository TemplateTagRepository
	StatusCode         int
}

func (interactor *TemplateInteractor) ListTemplates(limit int, offset int, keyword string) (t entity.Templates, totalPages int, err error) {
	t, err = interactor.TemplateRepository.Get(limit, offset, keyword)
	totalPages, err = interactor.TemplateRepository.Count()
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	interactor.StatusCode = 200
	return
}

func (interactor *TemplateInteractor) GetByUniqueId(uid string) (t entity.Template, tags []entity.Tag, err error) {
	t, err = interactor.TemplateRepository.GetByUniqueId(uid)
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	// 該当テンプレートの TagID 一覧を取得, TagIDをもとにタグのタイトルを取得
	var templateTags entity.TemplateTags
	templateTags, err = interactor.TemplateTagRepository.GetByTemplateId(t.Id)
	if err != nil {
		interactor.StatusCode = 500
		return
	}

	for _, tt := range templateTags {
		t, _ := interactor.TagRepository.GetById(tt.TagId)
		tags = append(tags, t)
	}
	fmt.Println(tags)

	interactor.StatusCode = 200
	return
}

func (interactor *TemplateInteractor) Add(template *entity.Template, tags []entity.Tag) (tp entity.Template, tg []entity.Tag, err error) {
	err = interactor.TemplateRepository.Insert(template)
	if err != nil {
		interactor.StatusCode = 500
		log.Fatalln(err)
		return
	}
	// TagをInsertしつつ、TemplateTagsにもInsert
	// err握り潰しちゃいます。
	for _, t := range tags {
		// Titleが重複するものもInsertしてしまいます
		requestTag := &entity.Tag{
			Title: t.Title,
		}
		_ = interactor.TagRepository.IgnoreInsert(requestTag)
		// TemplateTags
		tt := &entity.TemplateTag{
			TagId:      requestTag.Id,
			TemplateId: template.Id,
		}
		_ = interactor.TemplateTagRepository.Insert(tt)
	}

	// 追加されたレコードを取得
	tp, err = interactor.TemplateRepository.GetByUniqueId(template.Uid)
	if err != nil {
		interactor.StatusCode = 500
		return
	}
	interactor.StatusCode = 201
	return
}
