package main

import (
	"go_learn/function"
	"fmt"
)

func main() {
	nums := []int{7, 2, 3, 4}

	fmt.Println(function.Sum1(nums...))
	fmt.Println(function.Sum2(len(nums), nums...))
	fmt.Println(function.Factorial(7))
}