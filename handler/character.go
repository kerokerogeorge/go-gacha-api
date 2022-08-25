package handler

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"time"

	"github.com/Songmu/flextime"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"

	database "github.com/kerokerogeorge/go-gacha-api/database"
	"github.com/kerokerogeorge/go-gacha-api/model"
)

type UpdateCharacterRequest struct {
	EmissionRate float64 `json:"emissionRate"`
}

type ResultCharacterResponse struct {
	Id   string `json:"userCharacterId"`
	ID   string `json:"characterId"`
	Name string `json:"name"`
}
type GachaRequest struct {
	Times int `json:"times"`
}

type GachaResultResponse struct {
	ID   string `json:"characterId"`
	Name string `json:"name"`
}

type CreateCharacterRequest struct {
	Name         string  `json:"name"`
	EmissionRate float64 `json:"emissionRate"`
}

type GachaListResponse struct {
	ID string `json:"gachaId"`
}

// ガチャ実行API
func GetCharacter(c *gin.Context) {
	var characters []model.Character
	var user model.User
	var req GachaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	if err := database.DB.Table("users").Where("token = ?", key).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		panic(err)
	}

	if err := database.DB.Find(&characters).Error; err != nil {
		panic(err)
	}

	var selectedCharacterId int
	results := []GachaResultResponse{}
	now := time.Now()
	time.Sleep(time.Second * 3)
	for i := 0; i < req.Times; i++ {
		selectedCharacterId = DrawGacha(characters)
		// numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれていれば、キャラクターIDをもとにキャラクターをDBから取得
		character := PickCharacter(selectedCharacterId)

		result := model.Result{UserId: user.ID, CharacterId: character.ID}
		db := database.DB.Table("user_characters").Create(&result)
		if db.Error != nil {
			panic(db.Error)
		}

		res := GachaResultResponse{ID: character.ID, Name: character.Name}
		results = append(results, res)
	}
	log.Println("経過: ", time.Since(now).Milliseconds())

	c.JSON(http.StatusOK, gin.H{"results": results})
}

func PickCharacter(selectedCharacterId int) model.Character {
	var character model.Character
	if err := database.DB.Table("characters").Where("id = ?", selectedCharacterId).First(&character).Error; err != nil {
		panic(err)
	}

	return character
}

func DrawGacha(characters []model.Character) int {
	// 1〜100の範囲でランダムに値を取得
	rand.Seed(time.Now().UnixNano())
	rand := float64(rand.Intn(100-1) + 1)

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

	// return strconv.Itoa(selectedCharacterId), nil
	return selectedCharacterId
}

// ユーザ所持キャラクター一覧取得
func GetCharacterList(c *gin.Context) {
	var user model.User
	var results []ResultCharacterResponse

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
		Scan(&results).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		panic(err)
	}

	sort.Slice(results, func(i, j int) bool { return results[i].Id < results[j].Id })

	c.JSON(http.StatusOK, gin.H{"characters": results})
}

func CreateGacha(c *gin.Context) {
	log.Println("START=============")
	log.Println("CALLED")
	newGacha, err := NewGacha()
	if err != nil {
		panic(err)
	}
	db := database.DB.Table("gachas").Create(&newGacha)
	if db.Error != nil {
		panic(db.Error)
	}
	log.Println(db)
	log.Println("END=============")
}

func GetGachaList(c *gin.Context) {
	var gachas []model.Gacha
	var res []GachaListResponse
	if err := database.DB.Find(&gachas).Error; err != nil {
		panic(err)
	}

	for _, gacha := range gachas {
		res = append(res, GachaListResponse{ID: gacha.ID})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
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

func NewULID() ulid.ULID {
	t := flextime.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy)
}

func NewGacha() (*model.Gacha, error) {
	now := flextime.Now()
	return &model.Gacha{
		ID:        NewULID().String(),
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
