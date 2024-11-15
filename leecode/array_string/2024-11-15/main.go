package main

import "math/rand"

func main() {

}

func hIndex(citations []int) int {

}

type RandomizedSet struct {
	data map[int]int
	nums []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		data: make(map[int]int, 0),
		nums: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.data[val]; !ok {
		this.nums[len(this.data)] = val
		this.data[val] = len(this.nums)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if v, ok := this.data[val]; ok {
		delete(this.data, val)
		this.nums = append(this.nums[:v], this.nums[v+1:]...)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.nums))]
}
