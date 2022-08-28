package repository

import "github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

// import (
// "context"
// "github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

// "github.com/gin-gonic/gin"
// )

type CharacterRepository interface {
	GetCharacters() ([]*model.Character, error)
	// FindBoardThreads(ctx context.Context, categories []string, nextDocumentId string, limit int, isOwn bool, userId string) ([]*model.BoardThread, string, error)
}
