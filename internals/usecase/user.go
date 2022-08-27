package usecase

import (
	"github.com/google/uuid"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type UserUsecase interface {
	Create(name string) (string, error)
	Get(token string) (*model.User, error)
	Update(user *model.User, name string) (*model.User, error)
	GetAll() ([]*model.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (uu *userUsecase) Create(name string) (string, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uu.userRepo.CreateUser(name, token.String())
}

func (uu *userUsecase) Get(token string) (*model.User, error) {
	return uu.userRepo.GetUser(token)
}

func (uu *userUsecase) Update(user *model.User, name string) (*model.User, error) {
	return uu.userRepo.UpdateUser(user, name)
}

func (uu *userUsecase) GetAll() ([]*model.User, error) {
	return uu.userRepo.GetUsers()
}
