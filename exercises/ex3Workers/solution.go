package ex3Workers

import (
	"fmt"
	"sync"
	"time"
)

func Worker[T any](jobs <-chan T, results chan<- T, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- job
	}
}

func ResultCollector[T any](results <-chan T, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range results {
		fmt.Println(result)
		time.Sleep(time.Second)
	}
}

func Dispatcher[T any](jobsAmount, workersAmount int, produce func(i int) T) {
	jobs := make(chan T, jobsAmount)
	results := make(chan T, jobsAmount)
	var wg sync.WaitGroup
	wg.Add(workersAmount)

	for i := 0; i < workersAmount; i++ {
		go Worker(jobs, results, &wg)
	}

	var resultWg sync.WaitGroup
	resultWg.Add(1)
	go ResultCollector(results, &resultWg)
	for i := 0; i < jobsAmount; i++ {
		jobs <- produce(i)
	}

	close(jobs)

	wg.Wait()

	close(results)

	resultWg.Wait()
}
