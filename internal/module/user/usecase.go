package user

import "go-api/internal/entity"

type UserUsecase interface {
	CreateUser(user *entity.User) error
	GetUsers() ([]entity.User, error)
}

type userUsecase struct {
	repo UserRepository
}

func NewUserUsecase(r UserRepository) UserUsecase {
	return &userUsecase{r}
}

func (u *userUsecase) CreateUser(user *entity.User) error {
	// business logic ใส่ตรงนี้
	return u.repo.Create(user)
}

func (u *userUsecase) GetUsers() ([]entity.User, error) {
	return u.repo.FindAll()
}