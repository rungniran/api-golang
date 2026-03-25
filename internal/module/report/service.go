package report

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-api/internal/errs"
	"go-api/internal/pkg/response"
)

func (h *Handler) GetAll(c *gin.Context) {
	data, err := h.usecase.GetAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error")
		return
	}
	response.Success(c, data)
}

type createReq struct {
	Title string `json:"title"`
}

func (h *Handler) Create(c *gin.Context) {
	var req createReq

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	if req.Title == "" {
		response.Error(c, http.StatusBadRequest, "title is required")
		return
	}

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
	IDs    []string `json:"ids" binding:"required"`
	Reason string   `json:"reason"`
}

func (h *Handler) Approve(c *gin.Context) {
	var req updateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(req.IDs) == 0 {
		response.Error(c, http.StatusBadRequest, "ids is required")
		return
	}

	if err := h.usecase.Approve(req.IDs, req.Reason); err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error")
		return
	}

	response.Success(c, "approved")
}

func (h *Handler) Reject(c *gin.Context) {
	var req updateReq

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(req.IDs) == 0 {
		response.Error(c, http.StatusBadRequest, "ids is required")
		return
	}

	if err := h.usecase.Reject(req.IDs, req.Reason); err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error")
		return
	}

	response.Success(c, "rejected")
}