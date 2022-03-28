package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func ConnectDatabase() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:learngo@tcp(127.0.0.1:3306)/learngo"))
	if err != nil {
		log.Fatalln("数据库连接失败：", err)
	}
	fmt.Println("数据库连接成功")
	return conn
}
