package handler

import (
	"github.com/gin-gonic/gin"
)

func SetApiRoutes(
	e *gin.Engine,
	uh UserHandler,
) *gin.Engine {
	user := e.Group("/user")
	{
		user.GET("/", uh.GetOne)           // ユーザ情報取得API
		user.POST("/create", uh.Create)    // ユーザ情報作成API
		user.PUT("/update", uh.UpdateUser) // ユーザ情報更新API
		user.GET("/list", uh.GetUsers)     // 全ユーザーの取得
		user.DELETE("/", uh.DeleteUser)    // ユーザーの削除
	}
	// character := e.Group("/character")
	// {
	// 	character.GET("/", handler.GetCharacters)       // 全キャラクターを取得
	// 	character.POST("/", handler.CreateCharacter)    // 全キャラクターを作成
	// 	character.PUT("/:id", handler.UpdateCharacter) // キャラクターの排出率の変更
	// 	character.GET("/list", handler.GetCharacterList) // ユーザ所持キャラクター一覧取得API
	// 	character.GET("/emmition_rates", handler.GetEmmitionRate)
	// }

	// gacha := e.Group("/gacha")
	// {
	// 	gacha.GET("/", handler.GetGacha)   // 全キャラクターを取得
	// 	gacha.POST("/draw", handler.GetCharacter) // ガチャ実行API
	// 	gacha.POST("/create", handler.CreateGacha)
	// 	gacha.GET("/list", handler.GetGachaList)
	// }
	return e
}
