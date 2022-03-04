package algorithm

import "fmt"

// 二分查找找出目标的key 迭代法
func BinarySearch(array []int, target int) int {
	arrayLen := len(array)
	left, right := 0, arrayLen
	for i := 0; i < arrayLen; i++ {
		//middleIndex := low + (high-low)/2
		middleIndex := left + (right-left)>>1
		if array[middleIndex] == target {
			return middleIndex
		} else if array[middleIndex] > target {
			right = middleIndex - 1
		} else {
			left = middleIndex + 1
		}
	}
	return -1
}

// 二分查找找出目标的key -递归法
func BinarySearchV2(array []int, target int, left int, right int) int {
	//middleIndex := left + (right-left)/2
	middleIndex := left + (right-left)>>1
	fmt.Println(middleIndex)
	if array[middleIndex] == target {
		return middleIndex
	} else if array[middleIndex] > target {
		return BinarySearchV2(array, target, left, right-1)
	} else {
		return BinarySearchV2(array, target, left+1, right)
	}
	return -1
}
