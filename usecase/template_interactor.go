package usecase

import (
	"ca-zoooom/entity"
)

type TemplateInteractor struct {
	TemplateRepository    TemplateRepository
	TagRepository         TagRepository
	TemplateTagRepository TemplateTagRepository
	StatusCode            int
}

type TemplateResponse struct {
	Tags entity.Tags
}

func (interactor *TemplateInteractor) ListTemplates(limit int, offset int) (t entity.Templates, totalPages int, err error) {
	t, err = interactor.TemplateRepository.Get(limit, offset)
	totalPages, err = interactor.TemplateRepository.Count()
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	interactor.StatusCode = 200
	return
}

func (interactor *TemplateInteractor) GetByUniqueId(uid string) (t entity.Template, tags entity.Tags, err error) {
	t, err = interactor.TemplateRepository.GetByUniqueId(uid)
	if err != nil {
		interactor.StatusCode = 404
		return
	}
	// 該当テンプレートの TagID 一覧を取得, TagIDをもとにタグのタイトルを取得
	//var templateTags entity.TemplateTags
	//templateTags, err = interactor.TemplateTagRepository.GetByTemplateId(t.Id)
	//if err != nil {
	//	interactor.StatusCode = 500
	//	return
	//}
	//log.Println(templateTags)

	interactor.StatusCode = 200
	return
}

func (interactor *TemplateInteractor) Add(template *entity.Template) (t entity.Template, err error) {
	err = interactor.TemplateRepository.Insert(template)
	if err != nil {
		interactor.StatusCode = 500
		return
	}
	t, err = interactor.TemplateRepository.GetByUniqueId(template.Uid)
	if err != nil {
		interactor.StatusCode = 500
		return
	}
	interactor.StatusCode = 201
	return
}
