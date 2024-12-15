package handler

import (
	"dev-mgmt/internal/domain"
	"dev-mgmt/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	uc *usecase.DeviceUseCase
}

func NewDeviceHandler(uc *usecase.DeviceUseCase) *DeviceHandler {
	return &DeviceHandler{uc: uc}
}

func (h *DeviceHandler) RegisterDevice(c *gin.Context) {
	var device domain.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.RegisterDevice(c.Request.Context(), &device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Device registered successfully"})
}

func (h *DeviceHandler) AuthenticateDevice(c *gin.Context) {
	cert := c.GetHeader("X-Cert")
	if cert == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Certificate is required"})
		return
	}

	result, err := h.uc.AuthenticateDevice(c.Request.Context(), cert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.IsValid {
		c.JSON(http.StatusOK, gin.H{"message": "Device authenticated successfully"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"reason": result.Reason})
	}
}
