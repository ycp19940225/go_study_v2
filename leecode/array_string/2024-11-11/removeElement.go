package main

import "fmt"

func main() {

	//test := []int{3, 2, 2, 3}
	//test := []int{0, 1}
	test := []int{0, 1, 2, 3, 0, 4, 2, 2}
	//test := []int{0, 1, 3, 0, 4, 2, 2, 2}
	//test := []int{0, 1, 0, 4, 2, 2, 2, 2}
	//test := []int{0, 1, 0, 3, 4, 2, 2, 2}

	count := removeElement(test, 2)
	fmt.Printf("%+v %d", test, count)
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
