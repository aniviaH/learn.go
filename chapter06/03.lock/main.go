package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 1; i++ {
		//countDict()
		//countDictGoroutineSafe()
		//countDictForgetUnlock()
		countDictLockPrice()
	}
}

func countDict() {
	fmt.Println("开始数")
	counterWorker := 10000
	var totalCount int64 = 0
	var wg = sync.WaitGroup{}
	wg.Add(counterWorker)
	for p := 0; p < counterWorker; p++ {
		go func(p int) {
			defer wg.Done()
			//time.Sleep(1 * time.Second)
			//fmt.Println("正在统计第", p, "页, 当前获取totalCount=", totalCount)
			totalCount += 100
		}(p)
	}
	wg.Wait()
	fmt.Println("预计有：", 100*counterWorker, "字")
	fmt.Println("总共有：", totalCount, "字")
}

func countDictGoroutineSafe() {
	fmt.Println("开始数")
	const counterWorker int = 10000

	var totalCount int64 = 0
	totalCountLock := sync.Mutex{}

	var wg = sync.WaitGroup{}
	wg.Add(counterWorker)
	for p := 0; p < counterWorker; p++ {
		go func(p int) {
			defer wg.Done()
			//fmt.Println("正在统计第", p, "页, 当前获取totalCount=", totalCount)
			totalCountLock.Lock()
			defer totalCountLock.Unlock()
			totalCount += 100
			//totalCountLock.Unlock() // 也可以用defer来解锁
		}(p)
	}
	wg.Wait()
	fmt.Println("预计有：", 100*counterWorker, "字")
	fmt.Println("总共有：", totalCount, "字")
}

func countDictForgetUnlock() {
	fmt.Println("开始数")
	const counterWorker int = 5

	var totalCount int64 = 5
	totalCountLock := sync.Mutex{}

	var wg = sync.WaitGroup{}
	wg.Add(counterWorker)
	for p := 0; p < counterWorker; p++ {
		go func(p int) {
			defer wg.Done()
			//fmt.Println("正在统计第", p, "页, 当前获取totalCount=", totalCount)
			totalCountLock.Lock()
			totalCount += 100
			//totalCountLock.Unlock()
		}(p)
	}
	wg.Wait()
	fmt.Println("预计有：", 100*counterWorker, "字")
	fmt.Println("总共有：", totalCount, "字")
}

func countDictLockPrice() {
	fmt.Println("开始数")
	const counterWorker int = 10

	var totalCount int64 = 00
	totalCountLock := sync.Mutex{}

	var wg = sync.WaitGroup{}
	wg.Add(counterWorker)
	for p := 0; p < counterWorker; p++ {
		go func(p int) {
			defer wg.Done()
			//fmt.Println("正在统计第", p, "页, 当前获取totalCount=", totalCount)
			totalCountLock.Lock()
			totalCount += 100
			fmt.Println("p:", p)
			if p%3 == 0 {
				time.Sleep(3 * time.Second)
			}
			totalCountLock.Unlock()
		}(p)
	}
	wg.Wait()
	fmt.Println("预计有：", 100*counterWorker, "字")
	fmt.Println("总共有：", totalCount, "字")
}
