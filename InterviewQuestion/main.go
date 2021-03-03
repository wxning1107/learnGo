package main

import "fmt"

func main() {
	p := live()
	fmt.Println(p == nil)

	fmt.Println(recursiveClimbingStairs(10))
	fmt.Println(iterationClimbingStairs(10))
	fmt.Println(iterationClimbingStairs2(10))
	fmt.Println(mergeSort([]int{3, 5, 8, 1, 4, 6, 9, 2}))
}
