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

var facStore = &dbFactoryStore{}

type dbFactoryStore struct {
	store map[string]DBFactory
}

type Conn struct {
}
type DBFactory interface {
	GetConnection() *Conn
}

func initMySqlFac(connStr string) DBFactory {
	return &MySqlDBFactory{}
}

type MySqlDBFactory struct {
	once sync.Once
}

func (MySqlDBFactory) GetConnection() *Conn {
	once.Do(func() {
		initMySqlFac("")
	})
	// todo
	return nil
}

var counter int = 0
var counterOnce sync.Once

func main() {
	{
		for i := 0; i < 10; i++ {
			fmt.Println("第x次：", i)
			counterOnce.Do(func() {
				fmt.Println("初始化")
				counter++
			})
		}
		fmt.Println("最终结果：", counter)
	}
	/*{
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
	}*/

	{
		//connStr := "xxxxxxxx"

	}
}
