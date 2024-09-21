package main

import (
	"fmt"
	"os"
	"strconv"

	algorithms "github.com/Aervyon/education-algorithms/algorithms"
)

func main() {
	var input string
	fmt.Print("Please choose the number of an algorithm to execute. Your options are:\n1. Quicksort\nPlease choose: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	choice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input:", input)
		fmt.Println("Expected number.")
		os.Exit(1)
	}

	switch choice {
	case 1:
		algorithms.QuicksortExec()
	default:
		fmt.Println("That algorithm doesn't exist.")
		os.Exit(1)
	}
}
