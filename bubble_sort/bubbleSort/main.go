package main

import "fmt"

// sweep just swaps numbers in non-decreasing order
func sweep(numbers []int) {
	// setting N as the size of our list
	var N = len(numbers)

	// we know firstIndex & secondIndex need to start at 0
	// and 1, so we can initialize them with those values
	var firstIndex = 0
	var secondIndex = 1

	// just loop till secondIndex if smaller than N
	for secondIndex < N {
		var firstNumber = numbers[firstIndex]
		var secondNumber = numbers[secondIndex]

		// compare firstIndex & secondIndex and swap
		// if they are out of order
		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		// once swap is finished, increment indexes
		// so that it will start looking for next pair
		firstIndex++
		secondIndex++
	}
}

// bubbleSort sorts an array of unordered items
// into non-decreasing order
func bubbleSort(numbers []int) {
	var N = len(numbers)
	var i int
	for i = 0; i < N; i++ {
		sweep(numbers)
	}
}

func main() {
	var numbers = []int{5, 4, 2, 3, 1, 0}
	fmt.Println("Unsorted:", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted:", numbers)
}
