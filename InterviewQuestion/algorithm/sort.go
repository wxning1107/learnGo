package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{3, 5, 8, 1, 4, 6, 9, 2}))
}

// merge sort
// Time complexity: O(nlogn), Space Complexity: O(n)
func mergeSort(source []int) []int {
	if len(source) < 2 {
		return source
	}

	left := mergeSort(source[:len(source)/2])
	right := mergeSort(source[len(source)/2:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	res := make([]int, 0)
	for len(left) > i && len(right) > j {
		if left[i] > right[j] {
			res = append(res, right[j])
			j++
			continue
		}
		res = append(res, left[i])
		i++
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}
