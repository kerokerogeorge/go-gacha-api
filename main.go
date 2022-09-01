package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/infrastructure/datasource"
	"github.com/kerokerogeorge/go-gacha-api/internals/interface/handler"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gacha-API Docs
// @description ガチャアプリのAPI仕様書です
// @host localhost:8000
func main() {
	r := gin.Default()
	datasource.DbConnect()
	r = NewGin(r)
	r.Run(":8000")
}

func NewGin(e *gin.Engine) *gin.Engine {
	// datasource
	ur := datasource.NewUserRepository(datasource.DB)
	cr := datasource.NewCharacterRepository(datasource.DB)
	gr := datasource.NewGachaRepository(datasource.DB)
	cerr := datasource.NewCharacterEmmitionRateRepository(datasource.DB)
	rr := datasource.NewResultRepository(datasource.DB)

	// usecase
	uu := usecase.NewUserUsecase(ur, rr)
	cu := usecase.NewCharacterUsecase(cr, cerr)
	gu := usecase.NewGachaUsecase(gr, cr, cerr)

	// handler
	uh := handler.NewUserHandler(uu)
	ch := handler.NewCharacterHandler(cu)
	gh := handler.NewGachaHandler(gu, cu, uu)

	e = handler.SetApiRoutes(e, uh, ch, gh)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e
}
