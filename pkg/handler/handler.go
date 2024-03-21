package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	SrvIf serviceInterface
}

type serviceInterface interface {
	PerformWcReq(height int) (error, string)
}

func NewHandler(srv serviceInterface) *Handler {
	return &Handler{srv}
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

	err, bc := h.SrvIf.PerformWcReq(height)
	if err != nil {
		c.String(http.StatusNotFound, "")
		fmt.Println(err)
		return
	}

	c.String(http.StatusOK, bc)
}
