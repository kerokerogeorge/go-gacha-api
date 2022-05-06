package main

import (
	"log"

	"github.com/gin-gonic/gin"
	database "github.com/kerokerogeorge/go-test-prod/database"
	"github.com/kerokerogeorge/go-test-prod/handler"
)

func main() {
	log.Println("/* ===========START============ */")
	r := gin.Default()
	database.DbConnect()

	// 仕様書のユーザ関連API
	r.POST("/user/create", handler.CreateUser)  // ユーザ情報作成API
	r.GET("/user/get", handler.GetUser)         // ユーザ情報取得API
	r.PATCH("/user/update", handler.UpdateUser) // ユーザ情報更新API

	// 仕様書のガチャ関連API
	r.POST("/gacha/draw", handler.GetCharacter) // ガチャ実行API

	// 仕様書のキャラクター関連API
	r.GET("/character/list", handler.GetCharacterList) // ユーザ所持キャラクター一覧取得API

	log.Println("/* ===========開発用API============ */")
	// 開発用API
	r.GET("/user", handler.GetUsers)                   // 全ユーザーの取得
	r.DELETE("/user", handler.DeleteUser)              // ユーザーの削除
	r.GET("/character", handler.GetCharacters)         // 全キャラクターを取得
	r.PATCH("/character/:id", handler.UpdateCharacter) // キャラクターの排出率の変更

	r.Run(":8000")
}
