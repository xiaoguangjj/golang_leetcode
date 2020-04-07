/*
找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"


func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Printf("重复数字为：%d \n",findRepeatNumber(nums))
}

func findRepeatNumber(nums []int) int {
	md := make(map[int]int)
	for _, v := range nums {
		md[v]++
	}

	for k, v := range md {
		if v > 1 {
			return k
		}
	}
	return -1
}


