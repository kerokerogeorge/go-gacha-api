package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kerokerogeorge/go-gacha-api/internals/config"
	"github.com/kerokerogeorge/go-gacha-api/internals/infrastructure/datasource"
	"github.com/kerokerogeorge/go-gacha-api/internals/infrastructure/externals"
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

	config := cors.Config{
		// アクセス許可するオリジン
		AllowOrigins: []string{
			"http://localhost:3030",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{"Access-Control-Allow-Headers", "Content-Length", "Content-Type", "Authorization", "x-token"},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	r.Use(cors.New(config))
	r = NewGin(r)
	r.Run(":8000")
}

func NewGin(e *gin.Engine) *gin.Engine {
	// externals
	ethClient := config.NewEthClient()
	// datasource
	ur := datasource.NewUserRepository(datasource.DB)
	cr := datasource.NewCharacterRepository(datasource.DB)
	gr := datasource.NewGachaRepository(datasource.DB)
	cerr := datasource.NewCharacterEmmitionRateRepository(datasource.DB)
	ucr := datasource.NewUserCharacterRepository(datasource.DB)
	ecr := externals.NewEthereumRepository(ethClient)

	// usecase
	uu := usecase.NewUserUsecase(ur, ucr)
	cu := usecase.NewCharacterUsecase(cr, cerr, ucr)
	gu := usecase.NewGachaUsecase(gr, ur, ucr, cr, cerr)
	ucu := usecase.NewUserCharacterUsecase(ucr)
	ctu := usecase.NewContractUsecase(ecr)

	// handler
	uh := handler.NewUserHandler(uu)
	ch := handler.NewCharacterHandler(cu)
	gh := handler.NewGachaHandler(gu, cu, uu)
	uch := handler.NewUserCharacterHandler(ucu)
	cth := handler.NewContractHandler(ctu)

	e = handler.SetApiRoutes(e, uh, ch, gh, uch, cth)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return e
}
