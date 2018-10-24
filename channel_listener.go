package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func PlayChannelListener() {
	timeStart := time.Now()
	errChan := make(chan error, 3)
	defer close(errChan)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	var err error
	go func() {
		for e := range errChan {
			fmt.Println(e)
			if e != nil && err == nil {
				err = e
				cancel()
			}
		}

	}()

	chanCart := createCart(ctx, errChan)
	chanOrder := createOrder(ctx, errChan)
	chanPayment := createPayment(ctx, errChan)

	loadDone := make(chan bool)
	var cart, order, payment string
	go func() {
		cart, order, payment = <-chanCart, <-chanOrder, <-chanPayment
		close(loadDone)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("ctx done", ctx.Err())
	case <-loadDone:
	}

	fmt.Println(time.Since(timeStart).Seconds(), "finish")

	// just to make sure resource canceled
	fmt.Scanln()
}

func waitChannel2() bool {
	return true
}

func waitChannel(ch ...interface{}) <-chan bool {
	a := make(chan bool)
	close(a)
	return a
}

func createCart(ctx context.Context, errChan chan error) chan string {
	result := make(chan string)
	go func() {
		resp, err := db2.SelectContext(ctx, time.Second)

		if err != nil {
			errChan <- err
		}
		result <- resp.(string)
	}()
	return result
}

func createOrder(ctx context.Context, errChan chan error) chan string {
	result := make(chan string)
	go func() {
		resp, err := db2.SelectContext(ctx, time.Second*2)
		err = errors.New("resource order error nih")
		if err != nil {
			errChan <- err
		}
		result <- resp.(string)
	}()
	return result
}

func createPayment(ctx context.Context, errChan chan error) chan string {
	result := make(chan string)
	go func() {
		resp, err := db2.SelectContext(ctx, time.Second*3)
		if err != nil {
			errChan <- fmt.Errorf("resource payment %s", err.Error())
		}
		result <- resp.(string)
	}()
	return result
}

type DB2 struct{}

var db2 DB2

func (DB2) BeginTX() {
	fmt.Println("begin")
}

func (DB2) CommitTX() {
	fmt.Println("commit")
}

func (DB2) SelectContext(ctx context.Context, duration time.Duration) (result interface{}, err error) {
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
