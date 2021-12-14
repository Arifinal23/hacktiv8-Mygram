package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/arfan21/golang-mygram/helper"
	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		c.JSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, errors.New("the request is allowed for logged in")))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	bearerToken := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	id, err := helper.ParseJwt(bearerToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.NewResponse(http.StatusUnauthorized, nil, err))
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user_id", id)
	c.Next()
}
