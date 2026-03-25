package report

import "go-api/internal/entity"

type Usecase interface {
	GetAll() ([]entity.Report, error)
	Create(title string) error
	Approve(ids []string, reason string) error
	Reject(ids []string, reason string) error
}

type usecase struct {
	repo Repository
}

func NewUsecase(r Repository) Usecase {
	return &usecase{r}
}

func (u *usecase) GetAll() ([]entity.Report, error) {
	return u.repo.FindAll()
}

func (u *usecase) Create(title string) error {
	return u.repo.Create(&entity.Report{
		Title:  title,
		Status: entity.StatusPending,
	})
}

func (u *usecase) Approve(ids []string, reason string) error {
	return u.repo.UpdateStatus(ids, entity.StatusApproved, reason)
}

func (u *usecase) Reject(ids []string, reason string) error {
	return u.repo.UpdateStatus(ids, entity.StatusRejected, reason)
}