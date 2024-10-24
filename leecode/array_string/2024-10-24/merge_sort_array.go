package main

import "sort"

func main() {

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i+1] = nums2[i]
	}
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
