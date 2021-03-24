package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randMilliTime := rand.Intn(1000)
	fmt.Println(time.Millisecond * time.Duration(randMilliTime))
	fmt.Println(time.Millisecond*time.Duration(1000-randMilliTime) + time.Millisecond*time.Duration(randMilliTime))
	//sort.Sort()
	//sort.Ints()
	//math.Ceil()
	n := 10
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}

	fmt.Println(depth)
}
