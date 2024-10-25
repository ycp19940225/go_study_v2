package main

func main() {

}
func twoSum(nums []int, target int) []int {
	res := make([]int, 2, 2)
	for k, v := range nums {
		for k2, v2 := range nums {
			if v+v2 == target && k != k2 {
				res[0] = k
				res[1] = k2
			}
		}
	}
	return res
}
