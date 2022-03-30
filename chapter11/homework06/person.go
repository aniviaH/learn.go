package main

type PersonInfo struct {
	Id     int64   `json:"id,omitempty" gorm:"primaryKey;column:id"`
	Name   string  `json:"name,omitempty" gorm:"column:name"`
	Sex    string  `json:"sex,omitempty" gorm:"column:sex"`
	Tall   float32 `json:"tall,omitempty" gorm:"column:tall"`
	Weight float32 `json:"weight,omitempty" gorm:"column:weight"`
	Age    int64   `json:"age,omitempty" gorm:"column:age"`
	//Time     int     `json:"time,omitempty" gorm:"column:time"`
	IsDelete int64 `json:"is_delete,omitempty" gorm:"column:is_delete"`
}

func (*PersonInfo) TableName() string {
	return "circle"
}
