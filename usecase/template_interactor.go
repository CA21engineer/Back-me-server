package usecase

import (
	"ca-zoooom/entity"
	"log"
)

type TemplateInteractor struct {
	TemplateRepository    TemplateRepository
	TagRepository         TagRepository
	TemplateTagRepository TemplateTagRepository
	StatusCode            int
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

func (interactor *TemplateInteractor) GetByUniqueId(uid string) (template entity.Template, tags []entity.Tag, err error) {
	template, err = interactor.TemplateRepository.GetByUniqueId(uid)
	if err != nil {
		interactor.StatusCode = 404
		return
	}

	// 中間テーブルからとってくる
	var templateTags entity.TemplateTags
	templateTags, err = interactor.TemplateTagRepository.GetByTemplateId(template.Id)
	if err != nil {
		interactor.StatusCode = 500
		return
	}

	// []entity.Tag の形にまとめる
	for _, tt := range templateTags {
		var tag entity.Tag
		tag, _ = interactor.TagRepository.GetById(tt.TagId)
		tags = append(tags, tag)
	}

	interactor.StatusCode = 200
	return
}

func (interactor *TemplateInteractor) Add(tp *entity.Template, tg []entity.Tag) (template entity.Template, tags []entity.Tag, err error) {
	err = interactor.TemplateRepository.Insert(tp)
	if err != nil {
		interactor.StatusCode = 500
		log.Fatalln(err)
		return
	}
	// TagをInsertしつつ、TemplateTagsにもInsert
	// err握り潰しちゃいます...
	for _, t := range tg {
		var existTag entity.Tag
		existTag, _ = interactor.TagRepository.GetByTitle(t.Title)
		// 既にTitleが一致するものがあれば、TagへのInsertを行わない
		if existTag.Title != "" {
			tt := &entity.TemplateTag{
				TagId:      existTag.Id,
				TemplateId: tp.Id,
			}
			_ = interactor.TemplateTagRepository.Insert(tt)
		} else {
			requestTag := &entity.Tag{
				Title: t.Title,
			}
			_ = interactor.TagRepository.Insert(requestTag)

			tt := &entity.TemplateTag{
				TagId:      requestTag.Id,
				TemplateId: tp.Id,
			}
			_ = interactor.TemplateTagRepository.Insert(tt)
		}
	}

	// 追加されたレコードを取得
	template, err = interactor.TemplateRepository.GetById(tp.Id)
	if err != nil {
		interactor.StatusCode = 500
		log.Fatalln(err)
		return
	}

	// 中間テーブルからとってくる
	var templateTags entity.TemplateTags
	templateTags, err = interactor.TemplateTagRepository.GetByTemplateId(tp.Id)
	if err != nil {
		interactor.StatusCode = 500
		log.Fatalln(err)
		return
	}

	// []entity.Tag の形にまとめる
	for _, tt := range templateTags {
		var tag entity.Tag
		tag, _ = interactor.TagRepository.GetById(tt.TagId)
		tags = append(tags, tag)
	}

	interactor.StatusCode = 201
	return
}
