package user

import (
	"net/http"
    "go-api/internal/entity"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase UserUsecase
}

func NewUserHandler(r *gin.Engine, u UserUsecase) {
	h := &UserHandler{u}

	api := r.Group("/api")
	api.POST("/users", h.CreateUser)
	api.GET("/users", h.GetUsers)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.usecase.CreateUser(&user)
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, _ := h.usecase.GetUsers()
	c.JSON(http.StatusOK, users)
}
