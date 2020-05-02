package controllers

import (
	"ca-zoooom/interfaces/db"
	"ca-zoooom/usecase"
)

type TagController struct {
	Interactor usecase.TagInteractor
}

func NewTagController(sqlHandler db.SqlHandler) *TagController {
	return &TagController{
		Interactor: usecase.TagInteractor{
			TagRepository: &db.TagRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TagController) Index(c Context) {
	tags, err := controller.Interactor.ListTags()
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, H{"tags": tags})
}
