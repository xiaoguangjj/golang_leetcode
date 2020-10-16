/*
977. 有序数组的平方
难度：简单
给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。



示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]
示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]


提示：

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A 已按非递减顺序排序。

*/

package main

import (
	"fmt"
	"sort"
)

//学生成绩结构体
type StuScore struct {
	//姓名
	name  string
	//成绩
	score int
}


func sortedSquares(A []int) []int{
	type StuScores []StuScore


	B := []int{}
	for _,value := range A{
		B = append(B, value * value)
	}
	sort.Ints(B)

	return B
}

func main()  {
	var A = []int{-4, -1, 0, 3, 10}
	fmt.Printf("%v",sortedSquares(A))
}