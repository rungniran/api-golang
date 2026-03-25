package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-api/internal/app"
	"go-api/internal/entity"
	"go-api/internal/infrastructure"
)

func main() {
	r := gin.Default()
		// ✅ CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Angular dev
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	db := infrastructure.NewDB()
	db.AutoMigrate(&entity.User{}, &entity.Report{})
	app.Setup(r, db)
	r.Run(":8080")
}
