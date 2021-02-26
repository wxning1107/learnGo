package main

import "fmt"

func main() {
	p := live()
	fmt.Println(p == nil)

	fmt.Println(recursiveClimbingStairs(10))
	fmt.Println(iterationClimbingStairs(10))
}
