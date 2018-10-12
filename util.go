package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(i int) time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(i))
}

var counter int

func get(inp int) {
	for {
		time.Sleep(time.Millisecond * (random(100)))
		fmt.Println("go ", inp, counter)
	}
}

func set() {
	for {
		time.Sleep(time.Millisecond * (random(100)))
		counter++
	}
}

func worker(jobs chan func()) {
	for n := range jobs {
		n()
	}
}
