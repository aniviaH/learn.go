package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	s := &Store{
		Max: 10,
	}
	pCount, cCount := 50, 50
	for i := 0; i < pCount; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
				Producer{}.Produce(s)
			}
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() {
			for {
				time.Sleep(1 * time.Second)
				Consumer{}.Consume(s)
			}
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Println("库存", s.DataCount)
}

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex
}

type Producer struct {
}

func (p Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者看到库存满了，不生产")
		return
	}
	fmt.Println("生产者开始生产+1")
	s.DataCount++
}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者看到没有库存了，不消费")
		return
	}
	fmt.Println("消费者消费-1")
	s.DataCount--
}
