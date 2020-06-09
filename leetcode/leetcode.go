package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
	m := make(map[string]string)
	a, b := m["111"]
	fmt.Println(a, b)
}

func twoSum(nums []int, target int) (ans []int) {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		index, ok := m[target - nums[i]]
		if ok {
			ans = append(ans, index, i)
		} else {
			m[nums[i]] = i
		}
	}
	return ans
}