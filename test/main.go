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
}
