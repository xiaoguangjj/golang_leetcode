/*
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0
*/

package main

import "fmt"

//思路：二分查找
func searchInsert(nums []int, target int) int {
	if target < nums[0]{
		return 0
	}
	if target > nums[len(nums)-1]{
		return len(nums)
	}
	left := 0
	right := len(nums)

	for left < right{
		mid:= (left + right) / 2
		if nums[mid] < target{  // target 位于右侧
			left = mid + 1
		}else {				// target 位于左侧
			right = mid
		}
	}
	return left
}


func  main()  {
	//var nums [4] int
	nums := [] int {1, 3, 5, 7}
	var target int = 0
	fmt.Printf("target 的坐标为：%d\n",searchInsert(nums, target))
}