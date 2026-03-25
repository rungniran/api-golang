package report

import (

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase Usecase
}

func NewHandler(r *gin.Engine, u Usecase) {
	h := &Handler{u}

	api := r.Group("/api/reports")

	api.GET("", h.GetAll)
	api.POST("", h.Create)
	api.POST("/approve", h.Approve)
	api.POST("/reject", h.Reject)
}