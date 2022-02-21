package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			b := getRandRankBase()
			fmt.Println(b)
		}()
	}
	time.Sleep(1 * time.Second)
}
func getRandRankBase() int {
	rand.Seed(time.Now().UnixNano())
	base := rand.Intn(10)
	return base
}
