package main

import "fmt"

func main() {

	//test := []int{3, 2, 2, 3}
	//test := []int{0, 1}
	//test := []int{0, 1, 2, 3, 0, 4, 2, 2}

	//test := []int{0, 1, 3, 0, 4, 2, 2, 2}
	//test := []int{0, 1, 0, 4, 2, 2, 2, 2}
	//test := []int{0, 1, 0, 3, 4, 2, 2, 2}

	//count := removeElement(test, 2)
	//fmt.Printf("%+v %d", test, count)

	//test2 := []int{1, 2, 3, 4, 5, 6, 7}
	test2 := []int{-1, -100, 3, 99}
	rotate(test2, 2)
	fmt.Printf("%+v ", test2)
}

func rotate(nums []int, k int) {
	// 4 5 6
	length := len(nums)
	res := make([]int, 0, length)

	res = nums[length-k : length]
	res = append(res, nums[0:length-k]...)

	for i := 0; i < length; i++ {
		nums[i] = res[i]
	}

}

func maxProfit2(prices []int) int {

	res := 0
	length := len(prices)
	for i, v := range prices {
		for j := i + 1; j < length; j++ {
			if prices[j] > v && prices[j]-v > res {
				res = prices[j] - v
			}
		}
	}
	return res
}

func maxProfit(prices []int) int {

	res := 0
	length := len(prices)
	for i, v := range prices {
		for j := i + 1; j < length; j++ {
			if prices[j] > v && prices[j]-v > res {
				res = prices[j] - v
			}
		}
	}
	return res
}

func rotate2(nums []int, k int) {
	// 4 5 6
	length := len(nums)
	res := make([]int, 0, length)

	res = nums[length-k : length]
	res = append(res, nums[0:length-k]...)

	for i := 0; i < length; i++ {
		nums[i] = res[i]
	}

}

func removeElement(nums []int, val int) int {
	count := 1
	i := 0
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}
	for i <= (len(nums) - count) {
		if nums[i] == val {
			// 往前移动一位
			for j := i; j < len(nums)-count; j++ {
				nums[j] = nums[j+1]
			}
			count++
		} else {
			i++
		}
	}
	return len(nums) - count + 1
}

func remove2(nums []int, val int) int {
	k := 0
	flag := 0
	// 1 根据重复项个数决定前进步数
	for i := 1; i < len(nums); i++ {
		if nums[k] != nums[i] {
			k += flag
			nums[k] = nums[i]
			flag = 0
		}
		if flag < 2 {
			flag++
		}
	}
	return k + 1
}

func removeDuplicates(nums []int) int {
	k := 0
	flag := true
	// [0,0,1,1,1,1,2,3,3]
	// [0,0,1,1,2,1,2,3,3]
	// [0,0,1,1,2,3,2,3,3]
	// 1 根据重复项个数决定前进步数
	for i := 1; i < len(nums); i++ {
		// 负责交换
		if nums[k] != nums[i] {
			k++
			nums[k] = nums[i]
			flag = true
		} else { // 负责k的步进
			if flag {
				k++
				flag = false
				nums[k] = nums[i]
			}
		}
	}
	return k + 1
}
