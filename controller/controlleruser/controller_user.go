package controlleruser

import (
	"net/http"

	"github.com/arfan21/golang-mygram/helper"
	"github.com/arfan21/golang-mygram/model/modeluser"
	"github.com/arfan21/golang-mygram/service/serviceuser"
	"github.com/gin-gonic/gin"
)

type ControllerUser interface {
	Create(ctx *gin.Context)
	Login(ctx *gin.Context)
	Update(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

type controller struct {
	srv serviceuser.ServiceUser
}

func New(srv serviceuser.ServiceUser) ControllerUser {
	return &controller{srv}
}

// Create new user
// @Tags users
// @Summary Create new user
// @Description Create new user
// @Accept  json
// @Produce  json
// @Param data body modeluser.Request true "data"
// @Success 201 {object} helper.BaseResponse{data=modeluser.Response} "CREATED"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 409 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "data conflict, like email already exist"
// @Router /users/register [POST]
func (c *controller) Create(ctx *gin.Context) {
	data := new(modeluser.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := c.srv.Create(*data)

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusCreated, response, nil))
}

// Login user
// @Tags users
// @Summary Login user
// @Description Login user
// @Accept  json
// @Produce  json
// @Param data body modeluser.RequestLogin true "data"
// @Success 200 {object} helper.BaseResponse{data=modeluser.ResponseLogin} "OK"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 404 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Record not found"
// @Router /users/login [POST]
func (c *controller) Login(ctx *gin.Context) {
	data := new(modeluser.RequestLogin)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	response, err := c.srv.Login(*data)

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

// Update user
// @Tags users
// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param data body modeluser.ExampleRequestUpdate true "data"
// @Success 200 {object} helper.BaseResponse{data=modeluser.Response} "OK"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /users [PUT]
func (c *controller) Update(ctx *gin.Context) {
	data := new(modeluser.Request)

	if err := ctx.ShouldBind(data); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	id := ctx.MustGet("user_id")

	data.ID = id.(uint)

	response, err := c.srv.Update(*data)

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}

// DeleteByID user
// @Tags users
// @Summary Delete user
// @Description Delete user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Success 200 {object} helper.BaseResponse{data=modeluser.ExampleResponseDelete} "OK"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 404 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Not Found"
// @Failure 401 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /users [DELETE]
func (c *controller) DeleteByID(ctx *gin.Context) {
	id := ctx.MustGet("user_id")

	err := c.srv.DeleteByID(id.(uint))

	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, map[string]interface{}{"message": "your account has been successfully deleted"}, nil))
}
