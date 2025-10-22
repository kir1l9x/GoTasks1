package ex5ChanTimeout

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ChanelWriterReaderWithTimeout(n int) {
	ch := make(chan int, 1)
	timer, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(2)
	go ChanelWriterWithTimeout(timer, &wg, ch, func(i int) int { return i * i })
	go ChanelReaderWithTimeout(timer, &wg, ch)
	wg.Wait()
}

func ChanelWriterWithTimeout[T any](ctx context.Context, wg *sync.WaitGroup, output chan<- T, produceData func(i int) T) {
	defer wg.Done()
	defer close(output)
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case output <- produceData(i):
			i++
		}
	}
}

func ChanelReaderWithTimeout[T any](ctx context.Context, wg *sync.WaitGroup, input <-chan T) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-input:
			if !ok {
				return
			}
			fmt.Println(data)
		}
	}
}
