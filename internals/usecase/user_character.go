package usecase

import (
	"log"

	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type UserCharacterUsecase interface {
	UpdateStatus(key string, status model.CharacterStatus, userCharacterIDs []uint) ([]*model.UserCharacter, error)
}

type userCharacterUsecase struct {
	userCharcacterRepo repository.UserCharcacterRepository
}

func NewUserCharacterUsecase(ucr repository.UserCharcacterRepository) UserCharacterUsecase {
	return &userCharacterUsecase{
		userCharcacterRepo: ucr,
	}
}

func (ucu *userCharacterUsecase) UpdateStatus(key string, status model.CharacterStatus, userCharacterIDs []uint) ([]*model.UserCharacter, error) {
	var updatedUserCharacters []*model.UserCharacter
	for _, userCharacterID := range userCharacterIDs {
		userCharacter, err := ucu.userCharcacterRepo.GetOne(userCharacterID)
		if err != nil {
			return nil, err
		}

		var characterStatus model.CharacterStatus
		if status == model.CharacterStatusSuccess {
			characterStatus = model.CharacterStatusSuccess
		} else {
			characterStatus = model.CharacterStatusFailed
		}

		updatedUserCharacter, err := ucu.userCharcacterRepo.UpdateUsercharacter(userCharacter, characterStatus)
		if err != nil {
			return nil, err
		}
		log.Println(updatedUserCharacter)

		updatedUserCharacters = append(updatedUserCharacters, updatedUserCharacter)
	}

	return updatedUserCharacters, nil
}
