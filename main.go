package main

import (
	"fmt"
	"math"
)

func isPerfectSquare(x int) bool {
	s := int(math.Sqrt(float64(x)))
	return s*s == x
}

func isFibonacci(n int) bool {
	return isPerfectSquare(5*n*n + 4) || isPerfectSquare(5*n*n - 4)
}

func findFibonacciSequences(data []int) [][]int {
	var sequences [][]int
	var currentSequence []int

	for _, num := range data {
		if isFibonacci(num) {
			currentSequence = append(currentSequence, num)
		} else {
			if len(currentSequence) > 1 {
				sequences = append(sequences, currentSequence)
			}
			currentSequence = []int{}
		}
	}

	// Add the last sequence if it's valid
	if len(currentSequence) > 1 {
		sequences = append(sequences, currentSequence)
	}

	return sequences
}

func main() {
	data := []int{1, 2, 3, 4, 5, 8, 13, 21, 22, 34, 55, 89, 144, 233}
	fibonacciSequences := findFibonacciSequences(data)
	fmt.Println(fibonacciSequences)
}