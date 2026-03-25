package user

import (
	"go-api/internal/entity"

	"gorm.io/gorm"
)

// interface
type UserRepository interface {
	Create(user *entity.User) error
	FindAll() ([]entity.User, error)
}

// struct (implementation)
type userRepo struct {
	db *gorm.DB
}

// ✅ ตัวนี้แหละที่คุณขาด
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

// method
func (r *userRepo) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) FindAll() ([]entity.User, error) {
	var users []entity.User
	return users, r.db.Find(&users).Error
}
