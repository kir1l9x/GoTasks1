package ex2ConcurrentExponentiation

import (
	"sync"
	"time"
)

func ConcurrentExponentiation(nums []int) []int64 {
	resultNums := make([]int64, len(nums))
	wg := sync.WaitGroup{}
	for i, num := range nums {
		wg.Add(1)
		go func(num int, index int) {
			defer wg.Done()
			resultNums[index] = int64(num * num)
		}(num, i)
	}

	wg.Wait()

	return resultNums
}

func SimpleExponentiation(nums []int) []int64 {
	resultNums := make([]int64, len(nums))
	for i, num := range nums {
		resultNums[i] = int64(num * num)
	}

	return resultNums
}

func ConcurrentExponentiationForSmallSliceBenchOnly(nums []int) []int64 {
	resultNums := make([]int64, len(nums))
	wg := sync.WaitGroup{}
	for i, num := range nums {
		wg.Add(1)
		go func(num int, index int) {
			defer wg.Done()
			resultNums[index] = int64(num * num)
			time.Sleep(time.Second)
		}(num, i)
	}

	wg.Wait()

	return resultNums
}

func SimpleExponentiationForSmallSliceBenchOnly(nums []int) []int64 {
	resultNums := make([]int64, len(nums))
	for i, num := range nums {
		resultNums[i] = int64(num * num)
		time.Sleep(time.Second)
	}

	return resultNums
}
