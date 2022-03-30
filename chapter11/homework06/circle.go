package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type Circle struct {
	conn *gorm.DB
}

func (c *Circle) publish(info *PersonInfo) *gorm.DB {
	resp := c.conn.Create(info)
	if err := resp.Error; err != nil {
		fmt.Println("发布圈子失败：", err)
		return nil
	}
	fmt.Println("发布圈子成功")
	return resp
}

func (c *Circle) listByName(name string) []*PersonInfo {
	result := make([]*PersonInfo, 0)
	//resp := c.conn.Where(&PersonInfo{Name: name, IsDelete: 0}).Find(&result)
	resp := c.conn.Where("name = ?", name).Where("is_delete = ?", 0).Find(&result)
	if err := resp.Error; err != nil {
		fmt.Println("查找失败：", err)
		return nil
	}

	var data, _ = json.Marshal(result)
	fmt.Printf("%s列表：%+v\n", name, string(data))
	return result
}

func (c *Circle) listAll() []*PersonInfo {
	result := make([]*PersonInfo, 0)
	resp := c.conn.Where("is_delete = ?", 0).Find(&result)
	if err := resp.Error; err != nil {
		fmt.Println("查找失败：", err)
		return nil
	}

	var data, _ = json.Marshal(result)
	fmt.Printf("所有用户列表：%+v\n", string(data))
	return result
}

func (c *Circle) update(info *PersonInfo) *gorm.DB {
	// 更改年龄、身高、体重
	resp := c.conn.Model(info).Select("id", "age", "tall", "weight").Updates(info)
	if err := resp.Error; err != nil {
		fmt.Printf("更新%+v失败：", info)
		fmt.Println(err)
		return nil
	}
	fmt.Printf("更新%+v成功\n", info)
	return resp
}

func (c *Circle) delete(idToDelete int64) {
	p := &PersonInfo{
		Id:       idToDelete,
		IsDelete: 1,
	}
	resp := c.conn.Select("id", "is_delete").Updates(p)
	if err := resp.Error; err != nil {
		fmt.Printf("删除id: %d 失败：", idToDelete)
		fmt.Println(err)
		return
	}
	fmt.Printf("删除id: %d 成功\n", idToDelete)
}
