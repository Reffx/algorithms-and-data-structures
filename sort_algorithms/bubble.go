package main

import "fmt"

func main() {
	//Simple bubble sort
	var numbers []int = []int{5, 4, 3, 2, 1, 0}
	fmt.Println("Our list of numbers is:", numbers)
	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)
	//Improved bubble sort
	var numbersImproved []int = []int{1, 2, 3, 2, 4, 5}
	fmt.Println("(Improved) Our list of numbers is:", numbersImproved)
	bubbleSortImproved(numbersImproved)
	fmt.Println("After sorting:", numbersImproved)
}

func bubbleSort(numbers []int) {
	var N int = len(numbers)
	var i int
	for i = 0; i < N; i++ {
		sweep(numbers)
		fmt.Println("Sweep Nr", i+1, numbers)
	}
}

func sweep(numbers []int) {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < N {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		if firstNumber > secondNumber {
			//Swap the numbers
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		firstIndex++
		secondIndex++
	}
}

// Improvements
// 1. At the first sweep the highest number is already at the end, so we don't need to check the full array anymore
// 2. Check if the list is already sorted
func bubbleSortImproved(numbers []int) {
	var N int = len(numbers)
	var i int
	for i = 0; i < N; i++ {
		if !sweepImproved(numbers, i) { // I2 if statement
			return
		}
		fmt.Println("Sweep Nr", i+1, numbers)
	}
}

// return boolean true is a value was swapped
func sweepImproved(numbers []int, prevPasses int) bool { // I1,I2
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false // I2

	for secondIndex < (N - prevPasses) { // I1
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		if firstNumber > secondNumber {
			//Swap the numbers
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true // I2
		}

		firstIndex++
		secondIndex++
	}
	return didSwap //I2
}
