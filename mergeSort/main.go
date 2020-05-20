package main

import "fmt"

// O(nlogn)
func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	left := mergeSort(arr[0 : n/2])
	right := mergeSort(arr[n/2:])

	return merge(left, right)
}

// O(n)
func merge(left, right []int) []int {
	l, r := len(left), len(right)
	result := make([]int, 0)
	m, n := 0, 0

	for m < l && n < r {
		if left[m] > right[n] {
			result = append(result, right[n])
			n++
			continue
		}
		result = append(result, left[m])
		m++
	}

	result = append(result, left[m:]...)
	result = append(result, right[n:]...)

	return result
}

func main() {
	arr := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	result := mergeSort(arr)
	fmt.Println(result)
}
