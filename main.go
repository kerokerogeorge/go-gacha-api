package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/infrastructure/datasource"
	"github.com/kerokerogeorge/go-gacha-api/internals/interface/handler"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

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

	// usecase
	uu := usecase.NewUserUsecase(ur)
	cu := usecase.NewCharacterUsecase(cr)
	gu := usecase.NewGachaUsecase(gr)

	// handler
	uh := handler.NewUserHandler(uu)
	ch := handler.NewCharacterHandler(cu)
	gh := handler.NewGachaHandler(gu)

	e = handler.SetApiRoutes(e, uh, ch, gh)
	return e
}
