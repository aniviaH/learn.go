package main

import (
	"fmt"
	"sync"
)

type rank struct {
	standard []string
}

var globalRank = &rank{}
var globalRankInitialized bool = false
var globalRankInitializedLock sync.Mutex

var once sync.Once

func init() {
	globalRank.standard = []string{"Asia"}
}

func initGlobalRankStandard(standard []string, i int) {
	globalRankInitializedLock.Lock()
	defer globalRankInitializedLock.Unlock()

	if globalRankInitialized == true {
		fmt.Println("已初始化过，当前work", i)
		return
	}
	fmt.Println("第", i, "位work, 进行初始化")
	globalRank.standard = standard
	globalRankInitialized = true
}

func initGlobalRankStandardByOnce(standard []string, i int) {
	fmt.Println("第", i, "位work, 尝试进行初始化--")
	once.Do(func() {
		fmt.Println("第", i, "位work, 进行初始化")
		globalRank.standard = standard
	})
}

func main() {
	standard := []string{"asia"}
	for i := 0; i < 10; i++ {
		go func(i int) {
			//initGlobalRankStandard(standard, i)
			initGlobalRankStandardByOnce(standard, i)
		}(i)
	}
	//time.Sleep(1 * time.Second)
	fmt.Println(globalRank)
	fmt.Println(globalRankInitialized)
}
