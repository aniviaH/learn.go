package main

import (
	"fmt"
)

func main() {
	//var money interface{} = 10
	//var money interface{} = int64(10)
	var money interface{} = float32(10.0)
	//var money interface{} = "10"

	/*switch newmoney := money.(type) {
	case int:
		temMoney := newmoney + 3.0
		//temMoney := newmoney + 3.1
		fmt.Println("money是 int", newmoney+11, temMoney)
	case int64:
		temMoney := newmoney + 22
		fmt.Println("money是 int64", newmoney+22, temMoney)
	case float64:
		temMoney := newmoney + 33
		fmt.Println("money是 float64", newmoney+33, temMoney)
	case float32:
		temMoney := newmoney + 44
		fmt.Println("money是 float32", newmoney+44, temMoney)
	case string:
		temMoney := newmoney + "55"
		fmt.Println("money是 string", newmoney+"55", temMoney)
	default:
		fmt.Println("money是 未知类型")
	}*/

	switch money.(type) {
	case int, int16, int32, int64:
		fmt.Println("money是 整数")
	case float32, float64:
		fmt.Println("money是 小数")
	default:
		fmt.Println("money是 未知类型")
	}
	//sha256.New()

	/*
		分支表达式 switch case
		分支表达式可以完全替代if/else 条件表达式，并且分支表达式可以做不一样的事情。
		• 可随时终止分支表达式
		• 并入其他条件分支继续执行
		• 原生支持类型判断
		• 类型判断的同时赋值新的变量，且类型自动匹配
		• Case 可以是多条件的
	*/
}
