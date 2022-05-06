package handler

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/kerokerogeorge/go-test-prod/database"
	"github.com/kerokerogeorge/go-test-prod/model"
)

type UpdateCharacterRequest struct {
	EmissionRate float64 `json:"emissionRate"`
}

type ResultCharacterResponse struct {
	Id   string `json:"userCharacterId"`
	ID   string `json:"characterId"`
	Name string `json:"name"`
}

// ガチャ実行API
func GetCharacter(c *gin.Context) {
	var characters []model.Character
	var character model.Character
	var user model.User

	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	if err := database.DB.Table("users").Where("token = ?", key).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		panic(err)
	}

	if err := database.DB.Find(&characters).Error; err != nil {
		panic(err)
	}

	// 1〜100の範囲でランダムに値を取得
	rand := float64(rand.Intn(100-1) - 1)
	sum := 0
	// キャラクターの排出率を合計
	for _, v := range characters {
		sum += v.EmissionRate
	}
	multipleAmt := float64(100) / float64(sum)

	// 排出率の合計を100％に合わせて、キャラクターに定義されている排出率の数値に合わせて重みをつけ、配列に格納
	s := []float64{}
	for _, v := range characters {
		s = append(s, math.Round((float64(v.EmissionRate) * float64(multipleAmt))))
	}

	// 重みづけをした数値をnum=0から足していき、numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれているか検証
	num := float64(0)
	selectedCharacterId := 0
	for i, v := range s {
		if num < rand && rand <= num+v {
			selectedCharacterId = i + 1
			break
		} else {
			num += v
		}
	}

	// numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれていれば、キャラクターIDをもとにキャラクターをDBから取得
	if err := database.DB.Table("characters").Where("id = ?", strconv.Itoa(selectedCharacterId)).First(&character).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		panic(err)
	}

	result := model.Result{UserId: user.ID, CharacterId: character.ID}
	db := database.DB.Table("user_characters").Create(&result)
	if db.Error != nil {
		panic(db.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": character})
}

// ユーザ所持キャラクター一覧取得
func GetCharacterList(c *gin.Context) {
	var user model.User
	var result []ResultCharacterResponse

	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	if err := database.DB.Table("users").Where("token = ?", key).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		panic(err)
	}

	if err := database.DB.Table("users").Select("user_characters.id, characters.id, characters.name").
		Joins("INNER JOIN user_characters ON user_characters.user_id = ?", user.ID).
		Joins("INNER JOIN characters ON user_characters.character_id = characters.id").
		Where("users.id = ?", user.ID).
		Scan(&result).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		panic(err)
	}

	if len(result) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No results"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"characters": result})
}

// ============
// 以下開発用
// ============

// キャラクターの排出率の変更
func UpdateCharacter(c *gin.Context) {
	var character model.Character

	if err := database.DB.Table("characters").Where("id = ?", c.Param("id")).First(&character).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		return
	}

	var input UpdateCharacterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.DB.Model(&character).Updates(input)
	if db.Error != nil {
		panic(db.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": character})
}

// 全キャラクターを取得
func GetCharacters(c *gin.Context) {
	var characters []model.Character
	if err := database.DB.Find(&characters).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": characters})
}
