package arrayslice

import (
	"fmt"
	"math"
)

func Second_Smallest_Element1(input []int) int {
	min := math.MaxInt64
	second_min := math.MaxInt64

	n := len(input)
	for i := 0; i < n; i++ {
		if input[i] < min {
			min = input[i]
		}
	}

	fmt.Println("min", min)
	for i := 0; i < n; i++ {
		if input[i] != min && input[i] < second_min {
			second_min = input[i]
		}
	}
	return second_min
}

func SecondSmallestElement2(input []int) int {
	if len(input) < 2 {
		fmt.Println("Array must have at least 2 elements")
		return -1
	}

	min := math.MaxInt64
	second_min := math.MaxInt64

	// Single loop to find min and second min
	for _, num := range input {
		if num < min {
			second_min = min
			min = num
		} else if num > min && num < second_min {
			second_min = num
		}
	}

	if second_min == math.MaxInt64 {
		fmt.Println("No second smallest element found")
		return -1
	}

	return second_min
}
