package handler

import (
	"wc/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/getBookCountByAuthor/{id}", h.getBookCountByAuthor)
	router.GET("/getBookByAuthor/{id}", h.getBookByAuthor)
	return router
}
