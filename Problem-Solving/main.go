package main

import (
	"fmt"
	problem "problem-solving/Problem1"
)

func main() {
	fmt.Println("solve")
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(problem.FindDuplicate(nums))
}
