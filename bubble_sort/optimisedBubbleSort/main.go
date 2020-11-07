package main

import "fmt"

func sweep(numbers []int, prevPasses int) bool {
	var N = len(numbers)
	var firstIndex = 0
	var secondIndex = 1
	var didSwap = false

	for secondIndex < (N - prevPasses) {
		var firstNumber = numbers[firstIndex]
		var secondNumber = numbers[secondIndex]

		if firstNumber < secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}

func bubbleSort(numbers []int) {
	var N = len(numbers)
	var i int
	for i = 0; i < N; i++ {
		if !sweep(numbers, i) {
			return
		}

	}
}

func main() {
	var numbers = []int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12}
	fmt.Println("Unsorted:", numbers)

	bubbleSort(numbers)
	fmt.Println("Sorted:", numbers)
}
