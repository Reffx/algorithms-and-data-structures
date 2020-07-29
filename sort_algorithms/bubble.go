package main

import "fmt"

func main() {
	var numbers []int = []int{5, 4, 3, 2, 1, 0}
	fmt.Println("Our list of numbers is:", numbers)
	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)
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
