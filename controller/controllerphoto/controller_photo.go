package controllerphoto

import (
	"net/http"
	"strconv"

	"github.com/arfan21/golang-mygram/helper"
	"github.com/arfan21/golang-mygram/model/modelphoto"
	"github.com/arfan21/golang-mygram/service/servicephoto"
	"github.com/gin-gonic/gin"
)

type ControllerPhoto interface {
	Create(ctx *gin.Context)
	GetByUserID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	srv servicephoto.ServicePhoto
}

func New(srv servicephoto.ServicePhoto) ControllerPhoto {
	return &controller{srv: srv}
}

// Delete a photo
// @Tags photos
// @Summary Delete a photo
// @Description Delete a photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param photoID path int true "ID of the photo"
// @Success 200 {object} helper.BaseResponse "SUCCESS"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /photos/:photoID [DELETE]
func (c *controller) Delete(ctx *gin.Context) {
	paramKeyID := ctx.Param("photoID")
	photoID, _ := strconv.Atoi(paramKeyID)
	err := c.srv.Delete(photoID)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, "Your Photo has been successfully deleted", nil))
}

// Update a photo
// @Tags photos
// @Summary Update a photo
// @Description Update a photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param photoID path int true "ID of the photo"
// @Param data body modelphoto.Request true "data"
// @Success 200 {object} helper.BaseResponse{data=modelphoto.ResponseUpdate} "SUCCESS"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /photos/:photoID [PUT]
func (c *controller) Update(ctx *gin.Context) {
	data := new(modelphoto.Request)

	err := ctx.ShouldBind(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}
	photoParamID := ctx.Param("photoID")
	photoID, err := strconv.Atoi(photoParamID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}

	// must have user id from bearer
	userID := ctx.MustGet("user_id")
	data.UserID = userID.(uint)

	update, err := c.srv.Update(*data, photoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, update, nil))
}

// Create new photo
// @Tags photos
// @Summary Create new photo
// @Description Create new photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param data body modelphoto.Request true "data"
// @Success 201 {object} helper.BaseResponse{data=modelphoto.Response} "CREATED"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Router /photos [POST]
func (c *controller) Create(ctx *gin.Context) {
	data := new(modelphoto.Request)

	err := ctx.ShouldBind(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusUnprocessableEntity, nil, err))
		return
	}

	// must have user id from bearer
	userID := ctx.MustGet("user_id")
	data.UserID = userID.(uint)

	response, err := c.srv.Create(*data)
	if err != nil {
		ctx.JSON(helper.GetStatusCode(err), helper.NewResponse(helper.GetStatusCode(err), nil, err))
		return
	}

	ctx.JSON(http.StatusCreated, helper.NewResponse(http.StatusCreated, response, nil))
}

// GetByUserID a photo
// @Tags photos
// @Summary GetByUserID a photo
// @Description GetByUserID a photo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer + user token"
// @Param photoID path int true "ID of the photo"
// @Param data body modelphoto.Request true "data"
// @Success 200 {object} helper.BaseResponse{data=modelphoto.ResponseGet} "SUCCESS"
// @Failure 400 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Bad Request"
// @Failure 401 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Unauthorization"
// @Failure 404 {object} helper.BaseResponse{errors=helper.ExampleErrorResponse} "Not Found"
// @Router /photos [GET]
func (c *controller) GetByUserID(ctx *gin.Context) {
	response, err := c.srv.GetPhotos()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.NewResponse(http.StatusBadRequest, nil, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.NewResponse(http.StatusOK, response, nil))
}
