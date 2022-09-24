package controller

import (
	"assignment3/helpers"
	"assignment3/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetStatus(ctx *gin.Context)
}

type controller struct {
	service service.Service
}

func NewController(s service.Service) Controller {
	return &controller{
		service: s,
	}
}

func (c *controller) GetStatus(ctx *gin.Context) {
	result, err := c.service.GetData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to get order")
		return
	}
	response := helpers.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, response)
}
