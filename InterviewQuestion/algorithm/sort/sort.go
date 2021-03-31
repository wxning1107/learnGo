package main

import "fmt"

func main() {
	data := []int{3, 5, 8, 1, 4, 6, 9, 2}
	fmt.Println(mergeSort(data))
	quickSort(data, 0, 7)
	fmt.Println(data)
}

// merge sort
// Time complexity: O(nlogn), Space Complexity: O(n), 算法复杂度可以联想递归树, 树高是log以2为底n的对数, 每一层的消耗为n
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

// quick sort
func quickSort(data []int, p, r int) {
	if p >= r {
		return
	}

	pivot := data[r]
	i := p
	for j := p; j < r; j++ {
		if data[j] < pivot {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[r] = data[r], data[i]

	quickSort(data, p, i-1)
	quickSort(data, i+1, r)
}
