package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/pkg/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/blockchain/:height", h.getBlockchainHandler)
	return router
}

func (h *Handler) getBlockchainHandler(c *gin.Context) {
	var height int
	if _, err := fmt.Sscanf(c.Param("height"), "%d", &height); err != nil {
		panic(err)
	}
	//fmt.Printf("height: %d\n", height)

	err := h.service.PerformWcReq(height)
	if err != nil {
		c.String(http.StatusNotFound, "")
		fmt.Println(err)
		return
	}

	c.String(http.StatusOK, "")
}
