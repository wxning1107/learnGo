package main

import "fmt"

func main() {
	data := []int{3, 5, 7, 1, 2, 9, 4}
	quickSort(data, 0, len(data)-1)
	fmt.Println(data)
}

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
