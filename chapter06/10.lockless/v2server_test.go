package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWebServerV2(t *testing.T) {
	v1 := &WebServerV2{
		config: &Config{
			Content: "a",
		},
	}
	go v1.ReloadWorker()

	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				v1.Visit()
			}
		}()
	}
	wg.Wait()
	finish := time.Now()
	fmt.Println("总时间：", finish.Sub(start))
}
