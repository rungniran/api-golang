package infrastructure

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
