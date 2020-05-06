package controllers

import (
	"ca-zoooom/entity"
	"ca-zoooom/interfaces/db"
	"ca-zoooom/usecase"
	"math"
	"strconv"
)

type ImageController struct {
	Interactor usecase.ImageInteractor
}

type Pagination struct {
	CurrentPageNum int `json:"current_page_num"`
	ContentCount   int `json:"content_count"`
	TotalPage      int `json:"total_pages"`
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
	// ページネーション処理
	pageNumber, _ := strconv.Atoi(c.DefaultQuery("pages", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := limit * (pageNumber - 1)

	images, total, err := controller.Interactor.ListImages(limit, offset)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}

	// 全部のページ数。Goの言語仕様上、小数点切り上げを行うためにfloatで計算する必要がある
	totalPageCount := math.Ceil(float64(total) / float64(limit))

	c.JSON(controller.Interactor.StatusCode, H{"images": images, "pagination": Pagination{pageNumber, limit, int(totalPageCount)}})
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
	_ = c.Bind(&v)
	image, err := controller.Interactor.Add(v)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, image)
}

func (controller *ImageController) GetSignedUrl(c Context) {
	rowImgName := c.DefaultQuery("image", "default.jpeg")
	url, fileKey, err := controller.Interactor.GetSignedUrl(rowImgName)
	if err != nil {
		c.JSON(controller.Interactor.StatusCode, NewError(err))
		return
	}
	c.JSON(controller.Interactor.StatusCode, H{"signed_url": url, "file_key": fileKey})
}
