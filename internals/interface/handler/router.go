package handler

import (
	"github.com/gin-gonic/gin"
)

func SetApiRoutes(
	e *gin.Engine,
	uh UserHandler,
	ch CharacterHandler,
	gh GachaHandler,
) *gin.Engine {
	user := e.Group("/user")
	{
		user.GET("", uh.GetOne)            // ユーザ情報取得API
		user.POST("/create", uh.Create)    // ユーザ情報作成API
		user.PUT("/update", uh.UpdateUser) // ユーザ情報更新API
		user.GET("/list", uh.GetUsers)     // ユーザー一覧取得API
		user.DELETE("", uh.DeleteUser)     // ユーザー削除API
	}
	character := e.Group("/character")
	{
		character.GET("", ch.GetCharacters) // キャラクター一覧取得API
		character.POST("", ch.Create)       // 全キャラクターを作成API
		// // character.PUT("/:id", handler.UpdateCharacter) // キャラクターの排出率の変更
		// character.GET("/list", handler.GetCharacterList) // ユーザ所持キャラクター一覧取得API
		// character.GET("/emmition_rates", handler.GetEmmitionRate)
	}

	gacha := e.Group("/gacha")
	{
		gacha.POST("/create", gh.Create) // ガチャ作成API
		gacha.GET("/list", gh.List)      // キャラクター一覧取得API
		gacha.GET("", gh.Get)            // 一件のガチャ作成API
		gacha.POST("/draw", gh.Draw)     // ガチャ実行API
		gacha.DELETE("", gh.Delete)      // ガチャ削除API
	}
	return e
}
