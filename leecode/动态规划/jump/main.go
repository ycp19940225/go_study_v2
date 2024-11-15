package main

func main() {

}

func jump(nums []int) int {
	position := len(nums) - 1
	step := 0

	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				step++
				break
			}
		}
	}
	return step
}
