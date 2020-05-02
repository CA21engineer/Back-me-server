package controllers

import (
	"ca-zoooom/entity"
	"ca-zoooom/interfaces/db"
	"ca-zoooom/usecase"
	"strconv"
)

type VideoController struct {
	Interactor usecase.VideoInteractor
}

func NewVideoController(sqlHandler db.SqlHandler) *VideoController {
	return &VideoController{
		Interactor: usecase.VideoInteractor{
			VideoRepository: &db.VideoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *VideoController) Index(c Context) {
	videos, err := controller.Interactor.ListVideos()
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, H{"videos": videos})
}

func (controller *VideoController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	video, err := controller.Interactor.GetByID(id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, video)
}

func (controller *VideoController) Create(c Context) {
	v := &entity.Video{}
	c.Bind(&v)
	video, err := controller.Interactor.Add(v)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, video)
}

func (controller *VideoController) Update(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	v := &entity.Video{}
	c.Bind(&v)
	video, err := controller.Interactor.Update(v, id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, video)
}

func (controller *VideoController) Destroy(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := controller.Interactor.Delete(id)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, nil)
}
