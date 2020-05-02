package controllers

import (
	"ca-zoooom/entity"
	"ca-zoooom/interfaces/db"
	"ca-zoooom/usecase"
	"strconv"
)

type ImageController struct {
	Interactor usecase.ImageInteractor
}

func NewImageController(sqlHandler db.SqlHandler) *ImageController {
	return &ImageController{
		Interactor: usecase.ImageInteractor{
			ImageRepository: &db.ImageRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ImageController) Index(c Context) {
	images, err := controller.Interactor.ListImages()
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, H{"images": images})
}

func (controller *ImageController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	image, err := controller.Interactor.GetByID(id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, image)
}

func (controller *ImageController) Create(c Context) {
	v := &entity.Image{}
	c.Bind(&v)
	image, err := controller.Interactor.Add(v)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, image)
}
