package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	result := make(chan int, 200000)
	wg := sync.WaitGroup{}

	workNumber := 16
	baseNumCh := make(chan int, 1024)
	wg.Add(workNumber)
	for i := 0; i < workNumber; i++ {
		go func() {
			defer wg.Done()
			for oNum := range baseNumCh {
				if isPrime(oNum) {
					result <- oNum
				}
			}
		}()
	}

	for num := 2; num <= 200000; num++ {
		baseNumCh <- num
	}
	close(baseNumCh)
	wg.Wait()

	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
}

func isPrime(num int) (isPrime bool) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			return
		}
	}
	isPrime = true
	return
}
