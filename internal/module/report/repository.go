package report

import (
	"go-api/internal/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Report, error)
	Create(report *entity.Report) error
	UpdateStatus(ids []uint, status string, reason string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) FindAll() ([]entity.Report, error) {
	var reports []entity.Report
	return reports, r.db.Find(&reports).Error
}

func (r *repo) Create(report *entity.Report) error {
	return r.db.Create(report).Error
}

func (r *repo) UpdateStatus(ids []uint, status string, reason string) error {
	return r.db.Model(&entity.Report{}).
		Where("id IN ?", ids).
		Where("status = ?", entity.StatusPending).
		Updates(map[string]interface{}{
			"status": status,
			"reason": reason,
		}).Error
}