package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"time"
)

type User struct {
	Name string
	Base int
	Fat  float64
}
type Rank struct {
	usersMap map[int]*User
	rankLock sync.Mutex
}

const UserCount = 1000

var Base = 1.0
var FloatBase = 0.1

func main() {
	rank := &Rank{
		usersMap: map[int]*User{},
	}

	fmt.Println("开始初始化rank的用户", rank)
	wgForRegister := sync.WaitGroup{}
	wgForRegister.Add(UserCount)
	registerRankUsers(rank, &wgForRegister)
	wgForRegister.Wait()
	fmt.Println("rank用户初始化完成", rank)

	serverWg := sync.WaitGroup{}
	serverWg.Add(1)

	// 更新数据
	for i := 0; i < UserCount; i++ {
		go func(i int) {
			for {
				time.Sleep(2 * time.Second)
				updateUsers(rank, i)
			}
		}(i)
	}

	// 获取用户
	for i := 0; i < UserCount; i++ {
		go func(i int) {
			for {
				time.Sleep(2 * time.Second)
				getUserData(rank, i)
			}
		}(i)
	}

	//time.Sleep(3 * time.Second)
	serverWg.Wait()
	//RankServiceWithMutex(rank)
}

func registerRankUsers(rank *Rank, wg *sync.WaitGroup) {
	for i := 0; i < UserCount; i++ {
		go func(i int) {
			rank.rankLock.Lock()
			defer rank.rankLock.Unlock()
			defer wg.Done()

			rank.usersMap[i] = &User{
				Name: strconv.Itoa(i),
				Base: getRandRankBase(),
				Fat:  0,
			}
		}(i)
	}
}
func updateUsers(rank *Rank, name int) {
	defer rank.rankLock.Unlock()
	rank.rankLock.Lock()

	// 0：减 1：加
	symbol := [2]int{0, 1}
	index := symbol[getRandIndexOf0Or1()]
	var newFat float64
	if index == 0 {
		newFat = float64(rank.usersMap[name].Base) - getRandRankBaseFloat()
	} else {
		newFat = float64(rank.usersMap[name].Base) + getRandRankBaseFloat()
	}
	fmt.Println("更新用户：", name, "，体脂：", newFat)
	rank.usersMap[name].Fat = newFat
}

func getUserData(rank *Rank, name int) {
	userData := rank.usersMap[name]
	r := getUserRank(rank, userData, name)
	//fmt.Println("获取用户：", name, "，体脂: ", userData.Fat)
	fmt.Println("获取用户：", name, "，排名为: ", r)
}

type Users []*User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Fat < u[j].Fat
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
	//tmp:=b[i]
	//b[i] = b[j]
	//b[j]=tmp
}

func getUserRank(rank *Rank, u *User, name int) (res int) {
	users := Users{}
	for _, val := range rank.usersMap {
		users = append(users, val)
	}
	sort.Sort(sort.Reverse(users))

	fats := []float64{}
	for _, item := range users {
		fats = append(fats, item.Fat)
	}

	res = -1
	for i, item := range fats {
		if item == u.Fat {
			res = i
			break
		}
	}
	return res
}

func getRandRankBase() int {
	rand.Seed(time.Now().Unix())
	base := rand.Intn(10)
	return base
}
func getRandIndexOf20() int {
	rand.Seed(time.Now().Unix())
	base := rand.Intn(20)
	return base
}
func getRandIndexOf0Or1() int {
	rand.Seed(time.Now().Unix())
	base := rand.Intn(2)
	return base
}
func getRandRankBaseFloat() float64 {
	//FloatBase += 0.1
	var floats = [20]float64{
		0.01, 0.02, 0.03, 0.04, 0.05,
		0.06, 0.07, 0.08, 0.09,
		0.1, 0.11, 0.12, 0.13, 0.14, 0.15,
		0.16, 0.17, 0.18, 0.19, 0.20,
	}
	return floats[getRandIndexOf20()]
}
