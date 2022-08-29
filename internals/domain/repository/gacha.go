package repository

import (
	// "context"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	// "github.com/gin-gonic/gin"
)

type GachaRepository interface {
	CreateGacha(gacha *model.Gacha) (*model.Gacha, error)
	List() ([]*model.Gacha, error)
}
