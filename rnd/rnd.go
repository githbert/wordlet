package rnd

import (
	"math/rand"
	"time"
)

func GetNumber(min int, max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seed)

	return rnd.Intn(max-min+1) + min
}
