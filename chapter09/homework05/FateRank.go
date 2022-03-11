package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

type User struct {
	Index int
	Name  string
	//Age    int64
	//Sex    string
	//High   float64
	//Weight float64
	Fat        int64
	RegisterCh chan *User
}

func (u *User) registerToRank() {
	fmt.Println("用户进行注册-姓名：", u.Name, ", 体脂：", u.Fat)
	time.Sleep(50 * time.Millisecond)
	u.RegisterCh <- u
}
func (u *User) updateUserInfo() {

}
func (u *User) getRankList() {

}

type Rank struct {
	init      sync.Once
	usersMap  map[int]*User
	usersList []*User
}

func (r *Rank) RecieveRegisterOfUser(user *User) {
	r.usersMap[user.Index] = user

	//r.sortUsers()
	// 每次有人注册或者更新体脂排行榜，更新排行榜的排序列表
	r.sortUsersByBubble()
	//r.sortUsersByQuick()
}
func (r *Rank) sortUsersByBubble() {
	for _, v := range r.usersMap {
		r.usersList = append(r.usersList, v)
	}
	r.usersList = sortUsersByBubble(r.usersList)
}
func sortUsersByBubble(arr []*User) []*User {
	//fmt.Println("arr排序前--：")
	//for _, u := range arr {
	//	fmt.Printf("%v\t", u.Fat)
	//}
	//fmt.Println("")

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j].Fat > arr[j+1].Fat {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	//fmt.Println("")
	//fmt.Println("arr排序后：")
	//for _, u := range arr {
	//	fmt.Printf("%v\t", u.Fat)
	//}

	return arr
}
func (r *Rank) sortUsersByQuick() {

}
func (r *Rank) getRankList() {
	usersList := []*User{}
	for _, v := range r.usersMap {
		usersList = append(usersList, v)
	}
	usersList = sortUsersByBubble(usersList)

	fmt.Println("排名\t", "姓名\t", "体脂")
	for i, v := range usersList {
		fmt.Println(i+1, "\t", v.Name, "\t", v.Fat)
	}
}

const UserCount = 5

//func initUser() {
//	for i := 0; i < UserCount; i++ {
//		go func(i int) {
//			user := &User{}
//		}(i)
//	}
//}

func main() {
	registerCh := make(chan *User)

	rank := &Rank{
		usersMap: map[int]*User{},
	}
	rank.init.Do(func() {

	})

	// 用户注册
	fmt.Println("用户注册")
	for i := 0; i < UserCount; i++ {
		go func(i int) {
			fat := getRandFat()
			//fmt.Println(fat)
			user := &User{
				Index:      i,
				Name:       fmt.Sprintf("用户%d", i),
				Fat:        fat,
				RegisterCh: registerCh,
			}
			user.registerToRank()
		}(i)
	}

	registerSucCounter := 0
	//go func() {
	for user := range registerCh {
		fmt.Println("用户注册完成-姓名:", user.Name, ", 体脂:", user.Fat)
		rank.RecieveRegisterOfUser(user)
		registerSucCounter++

		if registerSucCounter == UserCount {
			close(registerCh)
		}
	}
	//}()

	//time.Sleep(2 * time.Second)
	fmt.Println("查看rank信息-----")
	for v, _ := range rank.usersMap {
		fmt.Printf("用户：%v\n", v)
	}
	fmt.Println("")

	// 用户查询体脂排行
	//go func() {
	rank.getRankList()
	//}()
	for i := 0; i < UserCount; i++ {

	}
}

func getRandFat() int64 {
	i, _ := rand.Int(rand.Reader, big.NewInt(3000))
	return i.Int64()
}
