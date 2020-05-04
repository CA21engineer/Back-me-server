package controllers

import (
	"ca-zoooom/entity"
	"ca-zoooom/interfaces/db"
	"ca-zoooom/usecase"
	"math"
	"strconv"
)

type TemplateController struct {
	Interactor usecase.TemplateInteractor
}

func NewTemplateController(sqlHandler db.SqlHandler) *TemplateController {
	return &TemplateController{
		Interactor: usecase.TemplateInteractor{
			TemplateRepository: &db.TemplateRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TemplateController) Index(c Context) {
	// ページネーション処理
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pages", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := limit * (pageNumber - 1)

	templates, total, err := controller.Interactor.ListTemplates(limit, offset)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}

	// 全部のページ数。Goの言語仕様上、小数点切り上げを行うためにfloatで計算する必要がある
	totalPageCount := math.Ceil(float64(total) / float64(limit))

	c.JSON(controller.Interactor.StatusCode, H{"templates": templates, "pagination": Pagination{pageNumber, limit, int(totalPageCount)}})
}

func (controller *TemplateController) Show(c Context) {
	uid := c.Param("uid")
	template, _, err := controller.Interactor.GetByUniqueId(uid)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}

	c.JSON(controller.Interactor.StatusCode, template)
}

func (controller *TemplateController) Create(c Context) {
	v := &entity.Template{}
	_ = c.Bind(&v)
	template, err := controller.Interactor.Add(v)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, template)
}
