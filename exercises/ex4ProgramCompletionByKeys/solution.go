package ex4ProgramCompletionByKeys

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func Worker[T any](ctx context.Context, jobs <-chan T, results chan<- T, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			select {
			case <-ctx.Done():
				return
			case results <- job:
			}
		}
	}
}

func ResultCollector[T any](ctx context.Context, results <-chan T, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case result, ok := <-results:
			if !ok {
				return
			}
			fmt.Println(result)
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second):
			}
		}

	}
}

func Dispatcher[T any](jobsAmount, workersAmount int, produce func(i int) T) {
	jobs := make(chan T, jobsAmount)
	results := make(chan T, jobsAmount)
	var wg sync.WaitGroup
	wg.Add(workersAmount)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	defer log.Println("program cancelled by SIGINT")

	for i := 0; i < workersAmount; i++ {
		go Worker(ctx, jobs, results, &wg)
	}

	var resultWg sync.WaitGroup
	resultWg.Add(1)
	go ResultCollector(ctx, results, &resultWg)
	for i := 0; i < jobsAmount; i++ {
		select {
		case <-ctx.Done():
			return
		case jobs <- produce(i):
		}
	}

	close(jobs)

	wg.Wait()

	close(results)

	resultWg.Wait()
}
