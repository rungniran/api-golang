package main

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/app"
	"go-api/internal/entity"
	"go-api/internal/infrastructure"
)

func main() {
	r := gin.Default()
	db := infrastructure.NewDB()
	db.AutoMigrate(&entity.User{}, &entity.Report{})
	app.Setup(r, db)
	r.Run(":8080")
}
