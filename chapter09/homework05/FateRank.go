package main

import "sync"

type User struct {
	Name string
	Base int
	Fat  float64
}
type Rank struct {
	usersMap map[int]*User
	rankLock sync.Mutex
}

func main() {

}
