package usecase

import (
	"errors"
	"log"
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
	"github.com/kerokerogeorge/go-gacha-api/internals/helper"
)

type GachaUsecase interface {
	Create() (*model.Gacha, error)
	List() ([]*model.Gacha, error)
	Get(gachaId string) (*model.Gacha, error)
	Draw(ctx *gin.Context, gachaId string, times int, key string) ([]*model.Result, error)
	DrawWithTransaction(ctx *gin.Context, gachaId string, times int, key string, from string, to string, contract string, amount *big.Int) ([]*model.Result, string, *types.Receipt, error)
	Delete(gachaId string) error
	GetGachaCharacters(gachaId string) ([]*model.CharacterEmmitionRate, error)
	DeleteGachaCharacters(gachaCharacters []*model.CharacterEmmitionRate) error
	ListHistory(ctx *gin.Context, key string) ([]*model.Result, error)
}

type gachaUsecase struct {
	gachaRepo                 repository.GachaRepository
	userRepo                  repository.UserRepository
	userCharcacterRepo        repository.UserCharcacterRepository
	characterRepo             repository.CharacterRepository
	characterEmmitionRateRepo repository.CharacterEmmitionRateRepository
	ethereumRepo              repository.EthereumRepository
}

func NewGachaUsecase(
	gr repository.GachaRepository,
	ur repository.UserRepository,
	ucr repository.UserCharcacterRepository,
	cr repository.CharacterRepository,
	cerr repository.CharacterEmmitionRateRepository,
	er repository.EthereumRepository,
) GachaUsecase {
	return &gachaUsecase{
		gachaRepo:                 gr,
		userRepo:                  ur,
		userCharcacterRepo:        ucr,
		characterRepo:             cr,
		characterEmmitionRateRepo: cerr,
		ethereumRepo:              er,
	}
}

func (gu *gachaUsecase) Create() (*model.Gacha, error) {
	characters, err := gu.characterRepo.GetCharacters()
	if err != nil {
		return nil, err
	}

	if len(characters) == 0 {
		return nil, errors.New("no characters")
	}

	newGacha, err := model.NewGacha()
	if err != nil {
		panic(err)
	}

	gacha, err := gu.gachaRepo.CreateGacha(newGacha)
	if err != nil {
		return nil, err
	}

	// 排出率をキャラクターごとに出す
	for _, character := range characters {
		characterWithEmmitionRate, err := model.NewCharacterEmmitionRate(gacha.ID, character.ID)
		if err != nil {
			return nil, err
		}

		err = gu.characterEmmitionRateRepo.SetEmmitionRate(characterWithEmmitionRate)
		if err != nil {
			return nil, err
		}
	}

	return gacha, nil
}

func (gu *gachaUsecase) List() ([]*model.Gacha, error) {
	return gu.gachaRepo.List()
}

func (gu *gachaUsecase) Get(gachaId string) (*model.Gacha, error) {
	gacha, err := gu.gachaRepo.GetOne(gachaId)
	if err != nil {
		return nil, errors.New("record not found")
	}

	charactersWithEmmitionRate, err := gu.characterEmmitionRateRepo.GetCharacterWithEmmitionRate(gacha.ID)
	if err != nil {
		return nil, errors.New("characters not found")
	}

	gachaWithCharacters := &model.Gacha{
		Characters: charactersWithEmmitionRate,
	}
	return gachaWithCharacters, err
}

func (gu *gachaUsecase) Draw(ctx *gin.Context, gachaId string, times int, key string) ([]*model.Result, error) {
	user, err := gu.userRepo.GetUser(key)
	if err != nil {
		return nil, errors.New("authentication failed")
	}

	charactersWithEmmitionRate, err := gu.characterEmmitionRateRepo.GetCharacterWithEmmitionRate(gachaId)
	if err != nil {
		return nil, errors.New("characters not found")
	}

	var results []*model.Result
	for i := 0; i < times; i++ {
		// 1〜100の範囲でランダムに値を取得
		rand := helper.NewRandomNumber()

		sum := 0
		// キャラクターの排出率を合計
		for _, v := range charactersWithEmmitionRate {
			sum += v.EmissionRate
		}
		multipleAmt := float64(100) / float64(sum)

		// 排出率の合計を100％に合わせて、キャラクターに定義されている排出率の数値に合わせて重みをつけ、配列に格納
		emmitionRates := []float64{}
		for _, character := range charactersWithEmmitionRate {
			emmitionRates = append(emmitionRates, (float64(character.EmissionRate) * float64(multipleAmt)))
		}

		// 重みづけをした数値をnum=0から足していき、numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれているか検証
		num := float64(0)
		var selectedCharacterId int
		var emissionRate float64
		for i, v := range emmitionRates {
			if num < rand && rand <= num+math.Round(v) {
				selectedCharacterId, _ = strconv.Atoi(charactersWithEmmitionRate[i].CharacterID)
				emissionRate = math.Round(v*100) / 100
				break
			} else {
				num += math.Round(v)
			}
		}

		character, err := gu.characterRepo.GetCharacter(selectedCharacterId)
		if err != nil {
			return nil, err
		}

		newUserCharacter, err := model.NewUserCharacter(user.ID, character.ID, character.ImgUrl, emissionRate)
		if err != nil {
			return nil, err
		}

		userCharacter, err := gu.userCharcacterRepo.CreateUserCharacter(newUserCharacter)
		if err != nil {
			return nil, err
		}

		// numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれていれば、キャラクターIDをもとにキャラクターをDBから取得
		res := &model.Result{ID: userCharacter.ID, CharacterId: character.ID, Name: character.Name, ImgUrl: character.ImgUrl, Status: userCharacter.Status, EmissionRate: emissionRate}
		results = append(results, res)
	}

	return results, nil
}

func (gu *gachaUsecase) Delete(gachaId string) error {
	gachaCharacters, err := gu.characterEmmitionRateRepo.GetGachaCharacters(gachaId)
	if err != nil {
		return errors.New("record not found")
	}

	err = gu.DeleteGachaCharacters(gachaCharacters)
	if err != nil {
		return errors.New("delete gacha characters failed")
	}

	gacha, err := gu.gachaRepo.GetOne(gachaId)
	if err != nil {
		return errors.New("record not found")
	}

	return gu.gachaRepo.DeleteGacha(gacha)
}

func (gu *gachaUsecase) GetGachaCharacters(gachaId string) ([]*model.CharacterEmmitionRate, error) {
	return gu.characterEmmitionRateRepo.GetGachaCharacters(gachaId)
}

func (gu *gachaUsecase) DeleteGachaCharacters(gachaCharacters []*model.CharacterEmmitionRate) error {
	for _, gachaCharacter := range gachaCharacters {
		err := gu.characterEmmitionRateRepo.DeleteGachaCharacter(gachaCharacter)
		if err != nil {
			return err
		}
	}

	return nil
}

func (gu *gachaUsecase) ListHistory(ctx *gin.Context, key string) ([]*model.Result, error) {
	user, err := gu.userRepo.GetUser(key)
	if err != nil {
		return nil, errors.New("authentication failed")
	}
	return gu.userCharcacterRepo.GetHistory(user.ID)
}

func (gu *gachaUsecase) DrawWithTransaction(ctx *gin.Context, gachaId string, times int, key string, from string, to string, contract string, amount *big.Int) ([]*model.Result, string, *types.Receipt, error) {
	user, err := gu.userRepo.GetUser(key)
	if err != nil {
		return nil, "", nil, err
	}

	log.Println("1")
	tx, receipt, err := gu.ethereumRepo.RawTransaction(ctx, from, to, contract, amount)
	if err != nil {
		return nil, "", receipt, err
	}

	charactersWithEmmitionRate, err := gu.characterEmmitionRateRepo.GetCharacterWithEmmitionRate(gachaId)
	if err != nil {
		return nil, "", nil, err
	}

	var results []*model.Result
	for i := 0; i < times; i++ {
		// 1〜100の範囲でランダムに値を取得
		rand := helper.NewRandomNumber()

		sum := 0
		// キャラクターの排出率を合計
		for _, v := range charactersWithEmmitionRate {
			sum += v.EmissionRate
		}
		multipleAmt := float64(100) / float64(sum)

		// 排出率の合計を100％に合わせて、キャラクターに定義されている排出率の数値に合わせて重みをつけ、配列に格納
		emmitionRates := []float64{}
		for _, character := range charactersWithEmmitionRate {
			emmitionRates = append(emmitionRates, (float64(character.EmissionRate) * float64(multipleAmt)))
		}

		// 重みづけをした数値をnum=0から足していき、numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれているか検証
		num := float64(0)
		var selectedCharacterId int
		var emissionRate float64
		for i, v := range emmitionRates {
			if num < rand && rand <= num+math.Round(v) {
				selectedCharacterId, _ = strconv.Atoi(charactersWithEmmitionRate[i].CharacterID)
				emissionRate = math.Round(v*100) / 100
				break
			} else {
				num += math.Round(v)
			}
		}

		character, err := gu.characterRepo.GetCharacter(selectedCharacterId)
		if err != nil {
			return nil, "", nil, err
		}

		newUserCharacter, err := model.NewUserCharacter(user.ID, character.ID, character.ImgUrl, emissionRate)
		if err != nil {
			return nil, "", nil, err
		}

		userCharacter, err := gu.userCharcacterRepo.CreateUserCharacter(newUserCharacter)
		if err != nil {
			return nil, "", nil, err
		}

		// numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれていれば、キャラクターIDをもとにキャラクターをDBから取得
		res := &model.Result{ID: userCharacter.ID, CharacterId: character.ID, Name: character.Name, ImgUrl: character.ImgUrl, Status: userCharacter.Status, EmissionRate: emissionRate}
		results = append(results, res)
	}

	return results, tx, receipt, nil
}
