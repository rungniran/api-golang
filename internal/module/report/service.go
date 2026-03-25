package report

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-api/internal/errs"
	"go-api/internal/pkg/response"
	"net/http"
)

func (h *Handler) GetAll(c *gin.Context) {
	data, _ := h.usecase.GetAll()
	c.JSON(200, data)
}

type createReq struct {
	Title string `json:"title"`
}

func (h *Handler) Create(c *gin.Context) {
	var req createReq

	// ✅ validate input
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request")
		return
	}
	
	if req.Title == "" {
		response.Error(c, http.StatusBadRequest, "title is required")
		return
	}
	// ✅ call usecase
	err := h.usecase.Create(req.Title)
	if err != nil {

		switch {
		case errors.Is(err, errs.ErrInvalidTitle):
			response.Error(c, http.StatusBadRequest, err.Error())

		default:
			response.Error(c, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response.Created(c, "created", nil)
}

type updateReq struct {
	IDs    []uint `json:"ids"`
	Reason string `json:"reason"`
}

func (h *Handler) Approve(c *gin.Context) {
	var req updateReq
	c.ShouldBindJSON(&req)
	h.usecase.Approve(req.IDs, req.Reason)
	c.JSON(http.StatusOK, gin.H{"message": "approved"})
}

func (h *Handler) Reject(c *gin.Context) {
	var req updateReq
	c.ShouldBindJSON(&req)
	h.usecase.Reject(req.IDs, req.Reason)
	c.JSON(http.StatusOK, gin.H{"message": "rejected"})
}
