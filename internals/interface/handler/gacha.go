package handler

import (
	// "math"

	// "math/rand"

	"net/http"

	// "sort"

	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

type GachaHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Draw(c *gin.Context)
}

type GetGachaResponse struct {
	GachaId   string                             `json:"gachaId"`
	Character []*model.CharacterWithEmmitionRate `json:"character"`
}

type GachaListResponse struct {
	ID string `json:"gachaId"`
}

type GetGachaRequest struct {
	GachaId string `form:"gachaId"`
}

type CreateGachaRequest struct {
	Times   int    `json:"times"`
	GachaID string `json:"gachaId"`
}

type GachaResultResponse struct {
	ID   string `json:"characterId"`
	Name string `json:"name"`
}
type gachaHandler struct {
	gachaUsecase     usecase.GachaUsecase
	characterUsecase usecase.CharacterUsecase
	userUsecase      usecase.UserUsecase
}

func NewGachaHandler(gu usecase.GachaUsecase, cu usecase.CharacterUsecase, uu usecase.UserUsecase) *gachaHandler {
	return &gachaHandler{
		gachaUsecase:     gu,
		characterUsecase: cu,
		userUsecase:      uu,
	}
}

func (gh *gachaHandler) Create(c *gin.Context) {
	gacha, err := gh.gachaUsecase.Create()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"gachaId": gacha.ID})
}

func (gh *gachaHandler) Get(c *gin.Context) {
	var req GetGachaRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gacha, err := gh.gachaUsecase.Get(req.GachaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	charactersWithEmmitionRate, err := gh.characterUsecase.GetCharactersWithEmmitionRate(gacha.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Characters not found"})
		return
	}

	getGachaResponse := &GetGachaResponse{
		GachaId:   gacha.ID,
		Character: charactersWithEmmitionRate,
	}

	c.JSON(http.StatusOK, gin.H{"data": getGachaResponse})
}

func (gh *gachaHandler) List(c *gin.Context) {
	var res []GachaListResponse
	gachas, err := gh.gachaUsecase.List()
	if err != nil {
		panic(err)
	}

	for _, gacha := range gachas {
		res = append(res, GachaListResponse{ID: gacha.ID})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (gh *gachaHandler) Draw(c *gin.Context) {
	var req CreateGachaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	user, err := gh.userUsecase.Get(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	charactersWithEmmitionRate, err := gh.characterUsecase.GetCharactersWithEmmitionRate(req.GachaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Characters not found"})
		return
	}

	var results []*GachaResultResponse
	for i := 0; i < req.Times; i++ {
		character, err := gh.gachaUsecase.Draw(charactersWithEmmitionRate, user.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "draw gacha failed"})
			return
		}
		// numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれていれば、キャラクターIDをもとにキャラクターをDBから取得
		res := &GachaResultResponse{ID: character.ID, Name: character.Name}
		results = append(results, res)
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

// func GetCharacter(c *gin.Context) {
// 	var user model.User
// 	var req GachaRequest

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	key := c.Request.Header.Get("x-token")
// 	if key == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
// 		return
// 	}

// 	if err := database.DB.Table("users").Where("token = ?", key).First(&user).Error; err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
// 		panic(err)
// 	}

// 	var charactersWithEmmitionRate []*Character
// 	if err := database.DB.Table("gachas").Select("character_emmition_rates.character_id, characters.name, character_emmition_rates.emission_rate").
// 		Joins("INNER JOIN character_emmition_rates ON character_emmition_rates.gacha_id = ?", req.GachaID).
// 		Joins("INNER JOIN characters ON character_emmition_rates.character_id = characters.id").
// 		Where("gachas.id = ?", req.GachaID).
// 		Scan(&charactersWithEmmitionRate).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
// 		panic(err)
// 	}

// 	var selectedCharacterId int
// 	results := []GachaResultResponse{}

/*
		時間計測のためのAPI
		now := time.Now()
	  time.Sleep(time.Second * 3)
	  log.Println("ガチャを実行するAPI: ", req.Times)
*/
// 	for i := 0; i < req.Times; i++ {
// 		selectedCharacterId = DrawGacha(charactersWithEmmitionRate)
// 		// numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれていれば、キャラクターIDをもとにキャラクターをDBから取得
// 		character := PickCharacter(selectedCharacterId)

// 		result := model.Result{UserId: user.ID, CharacterId: character.ID}
// 		db := database.DB.Table("user_characters").Create(&result)
// 		if db.Error != nil {
// 			panic(db.Error)
// 		}

// 		res := GachaResultResponse{ID: character.ID, Name: character.Name}
// 		results = append(results, res)
// 	}
// 	// log.Println("経過: ", time.Since(now).Milliseconds())

// 	c.JSON(http.StatusOK, gin.H{"results": results})
// }

// @@@
// func PickCharacter(selectedCharacterId int) model.Character {
// 	var character model.Character
// 	if err := database.DB.Table("characters").Where("id = ?", selectedCharacterId).First(&character).Error; err != nil {
// 		panic(err)
// 	}

// 	return character
// }

// func DrawGacha(characters []*Character) int {
// 	// 1〜100の範囲でランダムに値を取得
// 	rand.Seed(time.Now().UnixNano())
// 	rand := float64(rand.Intn(100-1) + 1)

// 	sum := 0
// 	// キャラクターの排出率を合計
// 	for _, v := range characters {
// 		sum += v.EmissionRate
// 	}
// 	multipleAmt := float64(100) / float64(sum)

// 	// 排出率の合計を100％に合わせて、キャラクターに定義されている排出率の数値に合わせて重みをつけ、配列に格納
// 	s := []float64{}
// 	for _, v := range characters {
// 		s = append(s, math.Round((float64(v.EmissionRate) * float64(multipleAmt))))
// 	}

// 	// 重みづけをした数値をnum=0から足していき、numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれているか検証
// 	num := float64(0)
// 	selectedCharacterId := 0
// 	for i, v := range s {
// 		if num < rand && rand <= num+v {
// 			selectedCharacterId = i + 1
// 			break
// 		} else {
// 			num += v
// 		}
// 	}

// 	return selectedCharacterId
// }

// @@@@
// func DeleteGacha(c *gin.Context) {
// 	var req DeleteGachaRequest
// 	var gacha model.Gacha

// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := database.DB.Table("gachas").Where("id = ?", req.GachaID).First(&gacha).Error; err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Record Not Found"})
// 		panic(err)
// 	}

// 	db := database.DB.Delete(&gacha)
// 	if db.Error != nil {
// 		panic(db.Error)
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted"})
// }
