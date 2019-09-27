package generator

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("BCDFGHJKLMNPQRSTVWXYZ")

func RandStringRunes(n int) string {
	value := make([]rune, n)
	for i := range value {
		value[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(value)
}
