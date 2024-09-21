package algorithms

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"slices"
)

func QuicksortExec() {
	// Accept input from the user
	fmt.Print("Enter your input (comma-separated numbers): ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Split the input string into an array of numbers
	inputArray := strings.Split(input, ",")
	numbers := make([]int, len(inputArray))
	for i, numStr := range inputArray {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			fmt.Println("Invalid input:", input)
			fmt.Println(err)
			os.Exit(1)
		}
		numbers[i] = num
	}

	sorted, comparisons, exchanges := quicksort(numbers)
	fmt.Println("Sorted:", sorted)
	fmt.Println("Comparisons:", comparisons)
	fmt.Println("Exchanges:", exchanges)

	// Print the result
}

func quicksort(input []int) ([]int, int, int) {
	if len(input) < 2 {
		return input, 0, 0
	}

	comparisons := 0
	exchanges := 0
	if len(input) == 2 {
		comparisons++
		if input[0] > input[1] { // Pseudo line: 3
			exchanges++
			return []int{input[1], input[0]}, comparisons, exchanges // Pseudo line: 4
		}
		return input, comparisons, exchanges
	}

	// we expect the array to be at least 3 elements long

	arr := input[:3]
	slices.Sort(arr)
	medianIndex := slices.Index(input, arr[1])

	// Rearrange the numbers so that the median is at the beginning
	input[0], input[medianIndex] = input[medianIndex], input[0] // Pseudo line: 5
	exchanges++

	pivot := input[0]
	var lowerLimit, upperLimit int
	lowerLimit = 1
	upperLimit = len(input) - 1
	for lowerLimit <= upperLimit { // Pseudo line: 9
		comparisons++
		for input[lowerLimit] < pivot { // Pseudo line: 10
			comparisons++
			lowerLimit++ // Pseudo line: 11
		}
		comparisons++
		for input[upperLimit] > pivot { // Pseudo line: 12
			comparisons++
			upperLimit-- // Pseudo line: 13
		}
		if lowerLimit <= upperLimit { // Pseudo line: 14
			exchanges++
			input[lowerLimit], input[upperLimit] = input[upperLimit], input[lowerLimit] // Pseudo line: 15
		}
	}
	// exchange elements at pivot and upperLimit
	exchanges++
	input[0], input[upperLimit] = input[upperLimit], input[0] // Pseudo line: 16

	less, comp, ex := quicksort(input[:upperLimit]) // Pseudo line: 17
	comparisons += comp
	exchanges += ex
	greater, comp, ex := quicksort(input[upperLimit+1:]) // Pseudo line: 18
	comparisons += comp
	exchanges += ex
	// fmt.Println(less)
	// fmt.Println(greater)

	return append(append(less, pivot), greater...), comparisons, exchanges
}
