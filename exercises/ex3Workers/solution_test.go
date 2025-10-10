package ex3Workers

import "testing"

func TestWorkerPool(t *testing.T) {
	Dispatcher(10000, 6, func(i int) int { return i })
}

func TestWorkerPool2(t *testing.T) {
	Dispatcher(5, 10, func(i int) int { return i })
}
