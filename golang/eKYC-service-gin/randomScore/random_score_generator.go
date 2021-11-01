package randomscore

import (
	"math/rand"
	"time"
)

func GenerateScore() int {
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)

	return (y1.Intn(100))
}
