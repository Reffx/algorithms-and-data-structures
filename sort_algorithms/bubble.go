import "fmt"

func main() {
	var numbers []int = []int{5,4,3,2,1,0} 
	fmt.Println("Our list of numbers is:", numbers)
	sweep(numbers)
}

func sweep(numbers []int){
	var N int = len(numbers)
	var firstIndex int = 0;
	var secondIndex int = 1

	for secondIndex < N {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		fmt.Println("Comparing the following:", firstNumber, secondNumber)

		firstIndex++
		secondIndex++
	}
}