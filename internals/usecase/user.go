package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type UserUsecase interface {
	Create(name string, address string) (string, error)
	Get(token string) (*model.User, error)
	List() ([]*model.User, error)
	Update(user *model.User, name string) (*model.User, error)
	Delete(token string) error
	GetUserCharacters(token string) ([]*model.Result, error)
}

type userUsecase struct {
	userRepo           repository.UserRepository
	userCharcacterRepo repository.UserCharcacterRepository
}

func NewUserUsecase(ur repository.UserRepository, rr repository.UserCharcacterRepository) UserUsecase {
	return &userUsecase{
		userRepo:           ur,
		userCharcacterRepo: rr,
	}
}

func (uu *userUsecase) Create(name string, address string) (string, error) {
	token, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uu.userRepo.CreateUser(name, token.String(), address)
}

func (uu *userUsecase) Get(token string) (*model.User, error) {
	return uu.userRepo.GetUser(token)
}

func (uu *userUsecase) Update(user *model.User, name string) (*model.User, error) {
	return uu.userRepo.UpdateUser(user, name)
}

func (uu *userUsecase) List() ([]*model.User, error) {
	return uu.userRepo.GetUsers()
}

func (uu *userUsecase) Delete(token string) error {
	user, err := uu.userRepo.GetUser(token)
	if err != nil {
		return errors.New("authentication failed")
	}
	return uu.userRepo.DeleteUser(user)
}

func (uu *userUsecase) GetUserCharacters(token string) ([]*model.Result, error) {
	user, err := uu.userRepo.GetUser(token)
	if err != nil {
		return nil, errors.New("authentication failed")
	}
	return uu.userCharcacterRepo.GetResults(user.ID)
}
