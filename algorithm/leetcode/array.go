package main

import "strings"

// 两个数组交集
func arrayIntersect(nums1, nums2 []int) []int {
	nums1M := make(map[int]int)
	res := make([]int, 0)
	for _, v := range nums1 {
		nums1M[v]++
	}

	for _, v := range nums2 {
		if c, ok := nums1M[v]; ok && c > 0 {
			nums1M[v]--
			res = append(res, v)
		}
	}

	return res
}

// 有序数组交集
func sortArrayIntersect(nums1, nums2 []int) []int {
	i, j := 0, 0
	res := make([]int, 0)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
			continue
		} else if nums1[i] < nums2[j] {
			i++
			continue
		} else if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		}
	}

	return res
}

func longestOfCommonPrefix(strs []string) string {
	res := ""
	k := 0
	prefix := strs[0]
	for _, str := range strs {
		for strings.Index(str, prefix) {

		}

	}
}
