// Sort a list of animals (strings) in alphabetical order
// ["dog", "cat", "alligator", "cheetah", "rat", "moose", "cow",
//	"mink", "porcupine", "dung beetle", "camel", "steer", "bat",
//	"hamster", "horse", "colt", "bald eagle", "frog", "rooster"]

package main

import "fmt"

func sweep(animals []string, prevPasses int) bool {
	N := len(animals)
	firstIndex := 0
	secondIndex := 1
	var didSwap = false

	for secondIndex < N {
		firstString := animals[firstIndex]
		secondString := animals[secondIndex]

		if greaterThan(firstString, secondString) == true {
			animals[firstIndex] = secondString
			animals[secondIndex] = firstString
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}

func greaterThan(a, b string) bool {
	if a > b {
		return true
	}
	return false
}

func bubbleSort(animals []string) {
	N := len(animals)
	var i int
	for i = 0; i < N; i++ {
		if !sweep(animals, i) {
			return
		}
	}
}

// entrypoint
func main() {
	var animals = []string{"dog", "cat", "alligator", "cheetah", "rat", "moose", "cow", "mink", "porcupine",
		"dung beetle", "camel", "steer", "bat", "hamster", "horse", "colt", "bald eagle", "frog", "rooster"}
	fmt.Println("Unsorted:", animals)

	bubbleSort(animals)
	fmt.Println("Sorted:", animals)
}
