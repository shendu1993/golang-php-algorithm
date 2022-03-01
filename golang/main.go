package main

import (
	"github.com/golang-php-algorithm/golang/array"
	"github.com/golang-php-algorithm/golang/test"
)

func main() {

	//https://leetcode.com/problems/two-sum/
	//var nums = []int{2,7,11,15}
	//test.Dd(array.TwoSum(nums, 9))
	//test.Dd(array.TwoSumV2(nums, 9))

	//https://leetcode.com/problems/median-of-two-sorted-arrays/
	var nums1 = []int{1, 3}
	var nums2 = []int{2}
	var nums3 = []int{1, 3}
	var nums4 = []int{2, 4}
	test.Dd(array.FindMedianSortedArrays(nums1, nums2))
	test.Dd(array.FindMedianSortedArrays(nums3, nums4))
}
