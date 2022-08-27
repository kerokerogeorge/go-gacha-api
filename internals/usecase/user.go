package usecase

import (
	"github.com/google/uuid"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type UserUsecase interface {
	Create(name string) (string, error)
	Get(token string) (string, error)
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

func (uu *userUsecase) Get(token string) (string, error) {
	return uu.userRepo.GetUser(token)
}
