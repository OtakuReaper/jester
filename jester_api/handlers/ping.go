package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jester_api/services"
)

type PingHandler struct {
	pingService services.PingService
}

func NewPingHandler(pingService services.PingService) *PingHandler {
	return &PingHandler{pingService: pingService}
}

func (h *PingHandler) Ping (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": h.pingService.Message(),
	})
}