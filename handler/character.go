package handler

import (
	"log"
	// "math"
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
	Name string `json:"name"`
}

type GachaListResponse struct {
	ID string `json:"gachaId"`
}

type GetEmmitionRateRequest struct {
	GachaID string `form:"gachaId"`
}

type DeleteGachaRequest struct {
	GachaID string `form:"gachaId"`
}

type GetGachaRequest struct {
	GachaID string `form:"gachaId"`
}

type GetGachaResponse struct {
	GachaID    string       `json:"gachaId"`
	Characters []*Character `json:"characters"`
}

type Character struct {
	CharacterID  string `json:"characterId"`
	Name         string `json:"name"`
	EmissionRate int    `json:"emissionRate"`
}

type CharacterEmmitionRateResponse struct {
	ID           string `json:"id"`
	CharacterID  string `json:"characterId"`
	Name         string `json:"name"`
	EmissionRate int    `json:"emissionRate"`
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

	// sum := 0
	// キャラクターの排出率を合計
	// for _, v := range characters {
	// 	sum += v.EmissionRate
	// }
	// multipleAmt := float64(100) / float64(sum)

	// 排出率の合計を100％に合わせて、キャラクターに定義されている排出率の数値に合わせて重みをつけ、配列に格納
	s := []float64{}
	// for _, v := range characters {
	// 	s = append(s, math.Round((float64(v.EmissionRate) * float64(multipleAmt))))
	// }

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
	newGacha, err := NewGacha()
	if err != nil {
		panic(err)
	}
	db := database.DB.Table("gachas").Create(&newGacha)
	if db.Error != nil {
		panic(db.Error)
	}

	// 排出率をキャラクターごとに出す
	// var characterEmmitionRate model.CharacterEmmitionRate
	var characters []model.Character
	if err := database.DB.Find(&characters).Error; err != nil {
		panic(err)
	}

	if len(characters) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no characters"})
		return
	}

	for _, character := range characters {
		rand.Seed(time.Now().UnixNano())
		emmitionRate := rand.Intn(100-1) + 1
		log.Println("character.ID: " + character.ID)
		characterEmmitionRate := model.CharacterEmmitionRate{GachaID: newGacha.ID, CharacterID: character.ID, EmissionRate: emmitionRate}
		db := database.DB.Table("character_emmition_rates").Create(&characterEmmitionRate)
		if db.Error != nil {
			panic(db.Error)
		}
	}

	log.Println("END=============")
	c.JSON(http.StatusOK, gin.H{"gachaId": newGacha.ID})
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

func GetGacha(c *gin.Context) {
	var req GetGachaRequest
	var gacha model.Gacha

	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Table("gachas").Where("id = ?", req.GachaID).First(&gacha).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		panic(err)
	}

	characters, err := ToCharacterModel(c, req.GachaID)
	if err != nil {
		panic(err)
	}

	log.Println("========")
	log.Println(characters)
	log.Println("========")
	// res := &GetGachaResponse{
	// 	GachaID:    req.GachaID,
	// 	Characters: characters,
	// }

	c.JSON(http.StatusOK, gin.H{"data": characters})
}

func ToCharacterModel(c *gin.Context, gachaId string) (*GetGachaResponse, error) {
	var character []*Character
	if err := database.DB.Table("gachas").Select("character_emmition_rates.character_id, characters.name, character_emmition_rates.emission_rate").
		Joins("INNER JOIN character_emmition_rates ON character_emmition_rates.gacha_id = ?", gachaId).
		Joins("INNER JOIN characters ON character_emmition_rates.character_id = characters.id").
		Where("gachas.id = ?", gachaId).
		Scan(&character).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		panic(err)
	}

	getGachaResponse := &GetGachaResponse{
		GachaID:    gachaId,
		Characters: character,
	}
	return getGachaResponse, nil
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

func CreateCharacter(c *gin.Context) {
	var req CreateCharacterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCharacter, err := NewCharacter(req.Name)
	if err != nil {
		panic(err)
	}
	db := database.DB.Table("characters").Create(&newCharacter)
	if db.Error != nil {
		panic(db.Error)
	}
	c.JSON(http.StatusOK, gin.H{"data": newCharacter})
}

func NewCharacter(name string) (*model.Character, error) {
	now := flextime.Now()
	return &model.Character{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func GetEmmitionRate(c *gin.Context) {
	var req GetEmmitionRateRequest
	var characterEmmitionRateResponse []CharacterEmmitionRateResponse

	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Table("gachas").Select("character_emmition_rates.id, character_emmition_rates.character_id, characters.name, character_emmition_rates.emission_rate").
		Joins("INNER JOIN character_emmition_rates ON character_emmition_rates.gacha_id = ?", req.GachaID).
		Joins("INNER JOIN characters ON character_emmition_rates.character_id = characters.id").
		Where("gachas.id = ?", req.GachaID).
		Scan(&characterEmmitionRateResponse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": characterEmmitionRateResponse})
}

func DeleteGacha(c *gin.Context) {
	var req DeleteGachaRequest
	var gacha model.Gacha

	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Table("gachas").Where("id = ?", req.GachaID).First(&gacha).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Record Not Found"})
		panic(err)
	}

	db := database.DB.Delete(&gacha)
	if db.Error != nil {
		panic(db.Error)
	}
	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted"})
}
