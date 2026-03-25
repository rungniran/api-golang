package app

import (
	"go-api/internal/module/report"
	"go-api/internal/module/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB) {
	// =====================
	// REPORT MODULE
	// =====================
	reportRepo := report.NewRepository(db)
	reportUsecase := report.NewUsecase(reportRepo)
	report.NewHandler(r, reportUsecase)
	// =====================
	// USER MODULE
	// =====================
	userRepo := user.NewUserRepository(db)
	userUsecase := user.NewUserUsecase(userRepo)
	user.NewUserHandler(r, userUsecase)
}
