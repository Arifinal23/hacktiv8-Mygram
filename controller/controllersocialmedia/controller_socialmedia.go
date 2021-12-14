package controllersocialmedia

import (
	"net/http"
	"strconv"

	"github.com/arfan21/golang-mygram/helper"
	"github.com/arfan21/golang-mygram/model/modelsocialmedia"
	"github.com/arfan21/golang-mygram/service/servicesocialmedia"
	"github.com/gin-gonic/gin"
)

type ControllerSocialMedia interface {
	Create(ctx *gin.Context)
	GetList(ctx *gin.Context)
	UpdateByID(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type controller struct {
	srv servicesocialmedia.ServiceSocialMedia
}

func New(srv servicesocialmedia.ServiceSocialMedia) ControllerSocialMedia {
	return &controller{srv}
}

// Create new social media
// @Tags socialmedias
// @Summary Create new social media
// @Description Create social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param data body modelsocialmedia.Request true "data"
// @Success 201 {object} helper.BaseResponse{data=modelsocialmedia.Response} "CREATED"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Router /socialmedias [POST]
func (c *controller) Create(ctx *gin.Context) {
	data := new(modelsocialmedia.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := c.srv.Create(*data)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, response, nil))
}

// Get all social media
// @Tags socialmedias
// @Summary Get all social media
// @Description Get all social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Success 200 {object} helper.BaseResponse{data=modelsocialmedia.ResponseListWrapper} "SUCCESS"
// @Router /socialmedias [GET]
func (c *controller) GetList(ctx *gin.Context) {

	response, err := c.srv.GetList()
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

// Update by id social media
// @Tags socialmedias
// @Summary Update by id social media
// @Description Update by id social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param socialmediaid path int true "ID of the social media"
// @Param data body modelsocialmedia.Request true "data"
// @Success 200 {object} helper.BaseResponse{data=modelsocialmedia.Response} "SUCCESS"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 404 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Record not found"
// @Router /socialmedias/:socialmediaid [PUT]
func (c *controller) UpdateByID(ctx *gin.Context) {
	data := new(modelsocialmedia.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	idString := ctx.Param("socialmediaid")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	data.ID = uint(id)

	response, err := c.srv.UpdateByID(*data)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

// Delete by id social media
// @Tags socialmedias
// @Summary Delete by id social media
// @Description Delete by id social media
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param socialmediaid path int true "ID of the social media"
// @Success 200 {object} helper.BaseResponse{data=modelsocialmedia.Response} "SUCCESS"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 404 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Record not found"
// @Router /socialmedias/:socialmediaid [DELETE]
func (c *controller) DeleteByID(ctx *gin.Context) {
	idString := ctx.Param("socialmediaid")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	err = c.srv.DeleteByID(uint(id))
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, map[string]interface{}{"message": "Your social media has been successfully deleted"}, nil))
}
