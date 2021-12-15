package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	/*
		课后作业
		1• 创建一个一维数组，反转它

		2• 用多维数组写一个日历表
		需要考虑每个月的天数不同
		需要考虑是平年还是闰年来专门处理二月
		• 【提升篇】日历按照一周的宽度显示（第一列为周一），且每个日期匹配到对应的列
	*/
	/*
		课后作业
		3• 创建一个一维整数切片，并用循环对它由从小到大排序
		4• 对一副新扑克牌打乱顺序
		5• 有一个包含中英文的切片，如果是英文的，转换它们的大小写
	*/

	//execPractice1()

	execPractice2()

	// 测试数组作为函数参数，是按引用还是按值
	/*
		arr := [4]int{1, 2, 3, 4}
		fmt.Printf("arr内存地址：%p\n", &arr)
		retArr := doubleArr(arr)
		fmt.Printf("%p\n", &arr)
		fmt.Println(&arr)
		fmt.Println(retArr)
	*/
}
func doubleArr(arr [4]int) (retArr [4]int) {
	fmt.Printf("arr内存地址：%p\n", &arr)
	for i, val := range arr {
		arr[i] = val * 2
	}
	return arr
}

func execPractice1() {
	arr := []int{1, 2, 3, 4, 9, 5, 6, 7}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7}
	reverseArr(arr)
	reverseArr(arr2)

	//reverseArr2(arr)
	//reverseArr2(arr2)
}
func reverseArr2(arr []int) {
	fmt.Println("反转前：", arr)
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		/*
				arr[i], arr[j] = arr[j], arr[i]
			=>
				temp1 := arr[j]
				temp2 := arr[i]
				arr[i] = temp1
				arr[j] = temp2
		*/
		i++
		j--
	}
	fmt.Println("反转后：", arr)
}

func reverseArr(arr []int) {
	fmt.Println("反转前：", arr)
	length := len(arr)
	mid := length / 2
	p := 0
	//for {
	//	if p >= mid {
	//		break
	//	}
	//
	//	tmp := arr[p]
	//	arr[p] = arr[length-p-1]
	//	arr[length-p-1] = tmp
	//
	//	p++
	//}

	for ; p < mid; p++ {
		//tmp := arr[p]
		//arr[p] = arr[length-p-1]
		//arr[length-p-1] = tmp

		arr[p], arr[length-p-1] = arr[length-p-1], arr[p]
	}

	fmt.Println("反转后：", arr)
}

func execPractice2() {
	year := getInputYear()
	month := getInputMonth()
	isYearLeap := getIsYearLeap(year)
	monthDaysMap := getMonthDaysMap(isYearLeap)
	//fmt.Println("monthDaysMap", monthDaysMap)

	daysCountInTheMonth := monthDaysMap[month]
	//fmt.Printf("%d年%d月的天数为：%d\n", year, month, daysCountInTheMonth)

	generateCalendar(year, month, daysCountInTheMonth)
}

func getMonthDaysMap(isYearLeap bool) (daysCountMapOfMonth map[int]int) {
	// 是否闰年
	var daysCountInFeb = 0
	if isYearLeap {
		daysCountInFeb = 29
	} else {
		daysCountInFeb = 28
	}
	daysCountMapOfMonth = map[int]int{
		1:  31,
		2:  0,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	daysCountMapOfMonth[2] = daysCountInFeb
	return
}

func getIsYearLeap(year int) (isLeapYear bool) {
	return (year%4 == 0) && (year%100 != 0)
}

func getInputMonth() (month int) {
	fmt.Print("月份：")
	fmt.Scanln(&month)
	return month
}

func getInputYear() (year int) {
	fmt.Print("年份：")
	fmt.Scanln(&year)
	return
}

func generateCalendar(year int, month int, daysCountInTheMonth int) {
	// 生成日历
	/*
			年份：2021 月份：3
		日	一	二	三	四	五	六
			1	2	3	4	5	6
		7	8	9	10	11	12	13
	*/
	const CalendarRow = 6
	const CalendarCol = 7
	const CalendarFlatArr42Length = CalendarRow * CalendarCol

	// 当前月开始的周下标
	weekDaysStrArr, firstDayIndexInFirstWeek := getFirstDayIndexInFirstWeek(year, month)
	//fmt.Println(firstDayIndexInFirstWeek)

	// 填充一个长度42、表示一个月的日期的数组，月份前后用默认值0
	calendarFlatArr42 := [CalendarFlatArr42Length]int{}
	calendarFlatArr42 = fillCalendarFlatArr42(firstDayIndexInFirstWeek, daysCountInTheMonth, calendarFlatArr42)
	//fmt.Println("------calendarFlatArr42------", calendarFlatArr42)

	// 遍历calendarFlatArr42，转成[6][7]string的日历字符串形式
	calendarMatrix := [CalendarRow][CalendarCol]string{
		//[7]int{},
	}
	calendarMatrix = fillCalendarMatrix(CalendarFlatArr42Length, CalendarCol, calendarFlatArr42, calendarMatrix)
	//fmt.Println("------calendarMatrix-------", calendarMatrix)

	// 打印日历
	printCalendarTopTool(year, month)
	printCalendarHeader(weekDaysStrArr)
	printCalendarBody(calendarMatrix)
}

func printCalendarHeader(weekDaysStrArr [7]string) {
	for _, weekday := range weekDaysStrArr {
		fmt.Print(weekday, "\t")
	}
	fmt.Println()
}

func printCalendarTopTool(year int, month int) {
	fmt.Printf("\t\t年份：%d  月份：%d\n", year, month)
}

func printCalendarBody(calendarMatrix [6][7]string) {
	for _, row := range calendarMatrix {
		for _, col := range row {
			fmt.Print(col, "\t")
		}
		fmt.Println()
	}
}

func fillCalendarMatrix(CalendarFlatArr42Length int, CalendarCol int, calendarFlatArr42 [42]int, calendarMatrix [6][7]string) (calendarMatrixFilled [6][7]string) {
	for i := 0; i < CalendarFlatArr42Length; i++ {
		row := i / CalendarCol
		col := i % CalendarCol
		day := calendarFlatArr42[i]
		var dayStr string

		if day != 0 {
			if calendarFlatArr42[i] < 10 {
				//dayStr = " " + string(val)
				dayStr = " " + strconv.Itoa(day)
			} else {
				//dayStr = string(val)
				dayStr = strconv.Itoa(day)
			}
		} else {
			dayStr = ""
		}
		calendarMatrix[row][col] = dayStr
	}
	return calendarMatrix
}

func fillCalendarFlatArr42(firstDayIndexInFirstWeek time.Weekday, daysCountInTheMonth int, calendarFlatArr42 [42]int) (calendarFlatArr42Filled [42]int) {
	counterDay := 0
	// time.Weekday转为int
	startDayIndex := int(firstDayIndexInFirstWeek)
	//fmt.Println("startDayIndex:", startDayIndex)
	for i := 0; i < len(calendarFlatArr42); i++ {
		fillNum := 0
		if i < startDayIndex {
			// 当前月1号之前-上一个月
			fillNum = 0
		}
		if i >= startDayIndex {
			// 当前月开始
			counterDay++
			fillNum = counterDay
		}
		if counterDay > daysCountInTheMonth {
			// 当前月之后-下一个月
			fillNum = 0
		}
		calendarFlatArr42[i] = fillNum
	}
	return calendarFlatArr42
}

func getFirstDayIndexInFirstWeek(year int, month int) ([7]string, time.Weekday) {
	dayFirst := 1
	monthDateTime := time.Date(year, time.Month(month), dayFirst, 0, 0, 0, 0, time.Local)
	// 星期的下标 [日, 一, 二, 三, 四, 五, 六]
	weekDaysStrArr := [7]string{"Sun", "Mon", "Tues", "Wed", "Thur", "Fri", "Sat"}
	// 月份第1天
	firstDayIndexInFirstWeek := monthDateTime.Weekday()
	//firstDayStr := weekDaysStrArr[firstDayIndexInFirstWeek]
	//fmt.Printf("%d年%d月%d日是 星期%s\n", year, month, dayFirst, firstDayStr)
	return weekDaysStrArr, firstDayIndexInFirstWeek
}
