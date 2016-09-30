package main

import (
	"math/rand"
	"time"
)

//GenerateRandom to generate a random number with seed and len
func random(min, max int) int {
	rand.Seed(time.Now().UnixNano() / int64(time.Millisecond))
	return rand.Intn(max-min) + min
}
