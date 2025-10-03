package ex2ConcurrentExponentiation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nums []int

func TestExps(t *testing.T) {
	nums = []int{1, 2, 3, 4, 5}

	expectedResult := []int64{1, 4, 9, 16, 25}
	actualResult := ConcurrentExponentiation(nums)

	assert.Equal(t, expectedResult, actualResult)
}

func TestConcurrentExponentiationWithValuesFromTask(t *testing.T) {
	nums = []int{2, 4, 6, 8, 10}

	expectedResult := []int64{4, 16, 36, 64, 100}
	actualResult := ConcurrentExponentiation(nums)

	assert.Equal(t, expectedResult, actualResult)
}

func BenchmarkConcurrentExponentiationForBigSliceAndSmallOperationTime(b *testing.B) {
	nums = make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcurrentExponentiation(nums)
	}
}

func BenchmarkSimpleExponentiationForBigSliceAndSmallOperationTime(b *testing.B) {
	nums = make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SimpleExponentiation(nums)
	}
}

func BenchmarkConcurrentExponentiationForSmallSliceAndBigOperationTime(b *testing.B) {
	nums = make([]int, 30)
	for i := 0; i < 30; i++ {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcurrentExponentiationForSmallSliceBenchOnly(nums)
	}
}

func BenchmarkSimpleExponentiationForSmallSliceAndBigOperationTime(b *testing.B) {
	nums = make([]int, 30)
	for i := 0; i < 30; i++ {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SimpleExponentiationForSmallSliceBenchOnly(nums)
	}
}
