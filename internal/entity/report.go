package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	Title     string    `json:"title"`
	Status    string    `json:"status"` // pending, approved, rejected
	Reason    string    `json:"reason"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *Report) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.NewString()
	return
}

const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusRejected = "rejected"
)