package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{10, 20},
		{30, 200},
		{400, 50},
		{30, 20},
	}
	cost := MinCost(matrix)
	fmt.Println("min cost=", cost)
}

func getMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func MinCost(matrix [][]int) int {
	sum := 0
	for _, el := range matrix {
		sum += getMin(el[0], el[1])
	}
	return sum
}
