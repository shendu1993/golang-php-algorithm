package array

import (
	"github.com/golang-php-algorithm/golang/test"
)

// https://leetcode.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	k := 0 //计算次数
	//初始化一个map
	m := make(map[int]int)
	//循环数组
	for i := 0; i < len(nums); i++ {
		//获取当前的值，达到目标需要的差值
		another := target - nums[i]
		//如果map中存在，则直接返回
		k++
		if _, ok := m[another]; ok {
			test.DdCycleTimes(k)
			return []int{m[another], i}
		}
		//记录到map中
		m[nums[i]] = i
	}
	return nil
}
func TwoSumV2(nums []int, target int) []int {
	k := 0 //计算次数
	//初始化一个map
	//循环数组
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			k++
			if nums[i]+nums[j] == target {
				test.DdCycleTimes(k)
				return []int{i, j}
			}
		}
	}
	return nil
}
