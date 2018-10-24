package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Mutex usahakan harus dibungkus 1 struct dengan yang ingin di mutex
type HitCount struct {
	Hit   int64
	Mutex sync.Mutex
}

func PlayMutex() {
	var state = make(map[int]int)
	var mtx sync.Mutex
	for r := 0; r < 100; r++ {
		go func() {
			for {
				mtx.Lock()
				fmt.Println(state[1] + 1)
				mtx.Unlock()
				time.Sleep(time.Millisecond * 10)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				val := rand.Intn(100)
				mtx.Lock()
				state[2] = val
				mtx.Unlock()
				time.Sleep(time.Millisecond * 10)
			}
		}()
	}
	fmt.Scanln()
	fmt.Println(state)
}

func PlayMutexInteger() {
	var counter int64
	var mtx sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10000)
	for w := 0; w < 10000; w++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 20; i++ {
				mtx.Lock()
				counter++
				mtx.Unlock()
				time.Sleep(time.Millisecond * 10)
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
