package handler

import (
	"github.com/gin-gonic/gin"
	docs "github.com/kerokerogeorge/go-gacha-api/docs"
)

func SetApiRoutes(
	e *gin.Engine,
	uh UserHandler,
	ch CharacterHandler,
	gh GachaHandler,
	uch UserCharacterHandler,
) *gin.Engine {
	docs.SwaggerInfo.BasePath = ""
	user := e.Group("/user")
	{
		user.GET("/list", uh.GetUsers)                // ユーザー一覧取得API
		user.POST("", uh.Create)                      // ユーザ情報作成API
		user.GET("", uh.GetOne)                       // ユーザ情報取得API
		user.PUT("", uh.UpdateUser)                   // ユーザ情報更新API
		user.DELETE("", uh.DeleteUser)                // ユーザー削除API
		user.GET("/characters", uh.GetUserCharacters) // ユーザ所持キャラクター一覧取得API
	}
	character := e.Group("/character")
	{
		character.GET("/list", ch.GetCharacters)                          // キャラクター一覧取得API
		character.POST("", ch.Create)                                     // 全キャラクターを作成API
		character.GET("/emmition_rates/:gachaId", ch.GetWithEmmitionRate) // ガチャの持つキャラクターを排出率とともに取得するAPI
		character.DELETE("/:characterId", ch.Delete)                      // キャラクター削除API
	}

	gacha := e.Group("/gacha")
	{
		gacha.GET("/list", gh.List)           // キャラクター一覧取得API
		gacha.POST("", gh.Create)             // ガチャ作成API
		gacha.GET("/:gachaId", gh.Get)        // 一件のガチャ取得API
		gacha.POST("/draw/:gachaId", gh.Draw) // ガチャ実行API
		gacha.DELETE("/:gachaId", gh.Delete)  // ガチャ削除API
	}
	userCharacter := e.Group("/user_character")
	{
		userCharacter.PUT("", uch.Update) // ユーザーのキャラクターのステータス変更API
	}
	// test := e.Group("/test")
	// {
	// 	test.POST("/payload", gh.CreateTokenTransferTransaction)
	// }
	return e
}
