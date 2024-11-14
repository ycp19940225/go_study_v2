package main

import "fmt"

func main() {

	//硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
	//输入: coins = [1, 2, 5], amount = 11，输出: 3  解释: 11 = 5 + 5 + 1
	//输入: coins = [2], amount = 3，输出: -1

	test := []int{2}

	res := getCoinsMax(test, 3)

	fmt.Println(res)
}

func getCoinsMax(nums []int, amount int) int {
	length := len(nums) - 1
	res := 0
	count := make([]int, 0)
	flag := false
	for length >= 0 {
		if res+nums[length] > amount {
			length--
			continue
		} else if res+nums[length] == amount {
			count = append(count, nums[length])
			flag = true
			break
		}
		res += nums[length]
		count = append(count, nums[length])
	}

	if flag {
		return len(count)
	}

	return -1
}
