package main

import (
	"fmt"
	"time"
)

func main() {
	db := ConnectDatabase()
	circle := &Circle{
		// db挂到circle上
		conn: db,
	}

	var nameUser1 = fmt.Sprintf("用户%d", 1)
	var nameUser2 = fmt.Sprintf("用户%d", 2)
	var nameUser3 = fmt.Sprintf("用户%d", 3)
	var user1 = &PersonInfo{
		Id:       1,
		Name:     nameUser1,
		Sex:      "男",
		Tall:     1.75,
		Weight:   65,
		Age:      25,
		Time:     time.Now(),
		IsDelete: 0,
	}
	var user2 = &PersonInfo{
		Id:       2,
		Name:     nameUser2,
		Sex:      "女",
		Tall:     1.65,
		Weight:   55,
		Age:      24,
		Time:     time.Now(),
		IsDelete: 0,
	}
	var user3 = &PersonInfo{
		Id:       3,
		Name:     nameUser3,
		Sex:      "男",
		Tall:     1.75,
		Weight:   70,
		Age:      28,
		Time:     time.Now(),
		IsDelete: 0,
	}

	// 发布圈子
	circle.publish(user1)
	circle.publish(user2)
	circle.publish(user3)

	// 查询所有用户状态
	circle.listAll()
	// 查询用户1状态
	circle.listByName(nameUser1)
	// 查询用户100状态
	circle.listByName(fmt.Sprintf("用户%d", 100))

	// 更新用户3的状态：年龄 身高 体重
	var user3New = &PersonInfo{
		Id:     3,
		Name:   nameUser3,
		Age:    29,
		Tall:   1.76,
		Weight: 75,
	}
	circle.update(user3New)

	// 查询用户3状态
	circle.listByName(nameUser3)

	// 删除用户2状态
	var idToDelete int64 = 2
	circle.delete(idToDelete)
	// 查询用户3状态
	circle.listByName(nameUser2)

	// 查询所有用户状态
	circle.listAll()
}
