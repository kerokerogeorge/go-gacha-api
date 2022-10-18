package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
)

type GachaRepository interface {
	CreateGacha(gacha *model.Gacha) (*model.Gacha, error)
	List() ([]*model.Gacha, error)
	GetOne(gachaId string) (*model.Gacha, error)
	DeleteGacha(gacha *model.Gacha) error
	TransferToken(ctx *gin.Context) (string, error)
}
