package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}

	res := summaryRanges(nums)

	fmt.Println(res)
}

func summaryRanges(nums []int) []string {

	// 如果 key+1 连续 放入temp 数组 继续
	// 如果不连续 放入temp 数组 append res 并且清空
	temp := make([]int, 0)
	var res []string
	for k, num := range nums {
		temp = append(temp, num)
		if k == (len(nums)-1) || num+1 == nums[k+1] {
			continue
		}
		length := len(temp)
		if length > 1 {
			res = append(res, fmt.Sprintf("%d->%d", temp[0], temp[length-1]))
		} else {
			res = append(res, fmt.Sprintf("%d", temp[0]))
		}
		temp = temp[:0]
	}
	return res
}
