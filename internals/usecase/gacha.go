package usecase

import (
	// "log"
	"errors"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type GachaUsecase interface {
	Create() (*model.Gacha, error)
	List() ([]*model.Gacha, error)
	Get(gachaId string) (*model.Gacha, error)
	Draw(charactersWithEmmitionRate []*model.CharacterWithEmmitionRate, userId string) (*model.Character, float64, error)
	Delete(gacha *model.Gacha) error
	GetGachaCharacters(gachaId string) ([]*model.CharacterEmmitionRate, error)
	DeleteGachaCharacters(gachaCharacters []*model.CharacterEmmitionRate) error
}

type gachaUsecase struct {
	gachaRepo                 repository.GachaRepository
	characterRepo             repository.CharacterRepository
	characterEmmitionRateRepo repository.CharacterEmmitionRateRepository
}

func NewGachaUsecase(gr repository.GachaRepository, cr repository.CharacterRepository, cerr repository.CharacterEmmitionRateRepository) GachaUsecase {
	return &gachaUsecase{
		gachaRepo:                 gr,
		characterRepo:             cr,
		characterEmmitionRateRepo: cerr,
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
	return gu.gachaRepo.GetOne(gachaId)
}

func (gu *gachaUsecase) Draw(charactersWithEmmitionRate []*model.CharacterWithEmmitionRate, userId string) (*model.Character, float64, error) {
	// 1〜100の範囲でランダムに値を取得
	rand.Seed(time.Now().UnixNano())
	rand := float64(rand.Intn(100-1) + 1)

	sum := 0
	// キャラクターの排出率を合計
	for _, v := range charactersWithEmmitionRate {
		sum += v.EmissionRate
	}
	multipleAmt := float64(100) / float64(sum)

	// 排出率の合計を100％に合わせて、キャラクターに定義されている排出率の数値に合わせて重みをつけ、配列に格納
	s := []float64{}
	for _, v := range charactersWithEmmitionRate {
		s = append(s, (float64(v.EmissionRate) * float64(multipleAmt)))
	}

	// 重みづけをした数値をnum=0から足していき、numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれているか検証
	num := float64(0)
	var selectedCharacterId int
	var emmitionRate float64
	for i, v := range s {
		if num < rand && rand <= num+math.Round(v) {
			selectedCharacterId, _ = strconv.Atoi(charactersWithEmmitionRate[i].CharacterID)
			emmitionRate = math.Round(v*100) / 100
			break
		} else {
			num += math.Round(v)
		}
	}

	character, err := gu.characterRepo.GetCharacter(selectedCharacterId)
	if err != nil {
		return nil, 0, err
	}

	newResult, err := model.NewUserCharacter(userId, character.ID)
	if err != nil {
		return nil, 0, err
	}

	err = gu.characterRepo.CreateUserCharacter(newResult)
	if err != nil {
		return nil, 0, err
	}

	err = gu.gachaRepo.TransferToken()
	if err != nil {
		return nil, 0, err
	}

	return character, emmitionRate, nil
}

func (gu *gachaUsecase) Delete(gacha *model.Gacha) error {
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
