package helper

import (
	"math/rand"
	"time"
)

// 1〜100の範囲でランダムに値を取得
func NewRandomNumber() float64 {
	var randomNumber float64
	rand.Seed(time.Now().UnixNano())
	randomNumber = float64(rand.Intn(100-1) + 1)
	return randomNumber
}
