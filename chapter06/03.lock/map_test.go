package main

import (
	"fmt"
	"testing"
	"time"
)

// 对map 数据的并发的读写会引起panic
func TestMap(t *testing.T) {
	m := map[int]int{}
	for i := 0; i < 100; i++ {
		go func() {
			for {
				v := m[i]
				m[i] = v + 1
				fmt.Println("i=", m[i])
			}
		}()
	}
	time.Sleep(10 * time.Second)
}
