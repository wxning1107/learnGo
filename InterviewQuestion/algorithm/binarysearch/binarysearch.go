package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 8))
}

func binarySearch(a []int, target int) int {
	low, high := 0, len(a)

	for low < high {
		mid := (low + high) >> 1
		if a[mid] < target {
			low = mid + 1
		} else if a[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
