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
func (u *User) getMyRank() {

}

type Rank struct {
	init      sync.Once
	usersMap  map[int]*User
	usersList []*User
}

func (rank *Rank) AcceptRegisterOfUser(user *User) {
	rank.usersMap[user.Index] = user

	// 每次有人注册或者更新体脂排行榜，更新排行榜的排序列表
	//r.sortUsers()
}

//func (r *Rank) sortUsers() {
//	for _, v := range r.usersMap {
//		r.usersList = append(r.usersList, v)
//	}
//	r.usersList = sortUsersByBubble(r.usersList)
//}
func (rank *Rank) sortUsersByBubbleSort(arr []*User) []*User {
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
func (rank *Rank) sortUsersByQuickSort(arr []*User, start int, end int) []*User {
	//fmt.Println("arr排序前--：")
	//for _, u := range arr {
	//	fmt.Printf("%v\t", u.Fat)
	//}
	//fmt.Println("")

	pivotIdx := (start + end) / 2
	pivotV := *arr[pivotIdx]
	l, r := start, end
	for l <= r {
		for (*arr[l]).Fat < pivotV.Fat {
			l++
		}
		for (*arr[r]).Fat > pivotV.Fat {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	if l == r {
		l++
		r--
	}
	if r > start {
		rank.sortUsersByQuickSort(arr, start, r)
	}
	if l < end {
		rank.sortUsersByQuickSort(arr, l, end)
	}

	//fmt.Println("")
	//fmt.Println("arr排序后：")
	//for _, u := range arr {
	//	fmt.Printf("%v\t", u.Fat)
	//}
	return arr
}
func (rank *Rank) getRankListByBubbleSort() []*User {
	fmt.Println("冒泡排序 获取体脂排行榜")

	var usersList []*User
	for _, v := range rank.usersMap {
		usersList = append(usersList, v)
	}
	usersList = rank.sortUsersByBubbleSort(usersList)

	return usersList
}
func (rank *Rank) getRankListByQuickSort() []*User {
	fmt.Println("快排 获取体脂排行榜")
	var usersList []*User
	for _, v := range rank.usersMap {
		usersList = append(usersList, v)
	}
	usersList = rank.sortUsersByQuickSort(usersList, 0, len(usersList)-1)

	return usersList
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
		rank.AcceptRegisterOfUser(user)
		registerSucCounter++

		if registerSucCounter == UserCount {
			close(registerCh)
		}
	}
	//}()

	//time.Sleep(2 * time.Second)
	fmt.Println("查看rank信息-----")
	for v := range rank.usersMap {
		fmt.Printf("用户：%v\n", v)
	}
	fmt.Println("")

	// 通过bubble排序查询排行榜总信息
	//usersList := rank.getRankListByBubbleSort()
	// 通过bubble排序查询排行榜总信息
	usersList := rank.getRankListByQuickSort()
	// 打印信息
	fmt.Println("排名\t", "姓名\t", "体脂")
	for i, v := range usersList {
		fmt.Println(i+1, "\t", v.Name, "\t", v.Fat)
	}

}

func getRandFat() int64 {
	i, _ := rand.Int(rand.Reader, big.NewInt(3000))
	return i.Int64()
}
