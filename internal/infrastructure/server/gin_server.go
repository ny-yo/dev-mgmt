package server

import (
	"dev-mgmt/internal/interface/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(deviceHandler *handler.DeviceHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/device/register", deviceHandler.RegisterDevice)
		api.POST("/device/authenticate", deviceHandler.AuthenticateDevice)
	}

	return r
}
