package main

import (
	"context"
	"fmt"
	"time"
)

func PlayContext() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	doneChan := make(chan bool)
	defer cancel()

	db.BeginTX()
	var err error
	go func() {
		err = CreateOrder(ctx)
		doneChan <- false
	}()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	case <-doneChan:
	}
	db.CommitTX()
	fmt.Println("finish")
}

func CreateOrder(ctx context.Context) error {
	for i := 0; i < 5; i++ {
		_, err := db.SelectContext(ctx, time.Second*1000)
		if err != nil {
			return err
		}
	}
	fmt.Println("finish select")
	OtherProccess()
	return nil
}

func OtherProccess() {
	<-time.Tick(time.Second * 5)
}

type DB struct{}

var db DB

func (DB) BeginTX() {
	fmt.Println("begin")
}

func (DB) CommitTX() {
	fmt.Println("commit")
}

func (DB) SelectContext(ctx context.Context, duration time.Duration) (result interface{}, err error) {
	doneChan := make(chan bool)
	result = ""
	go func() {
		<-time.Tick(duration)
		result = "selected 1 row"
		doneChan <- false
	}()

	select {
	case <-ctx.Done():
		return result, ctx.Err()
	case <-doneChan:
		return result, err
	}
}
