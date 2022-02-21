package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Store struct {
	init  sync.Once
	store chan int
	Max   int
}

//func (s *Store) Close() {
//
//}

func (s *Store) instrument() {
	s.init.Do(func() {
		s.store = make(chan int, s.Max)
	})
}

type Producer struct {
}

func (p Producer) Produce(s *Store) {
	fmt.Println("生产者开始生产+1")
	s.store <- rand.Int()
}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {
	fmt.Println("消费者消费-1", <-s.store)
}

func main() {
	s := &Store{
		Max: 10,
	}
	s.instrument()
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
}
