/*
83. 删除排序链表中的重复元素
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/
package main

import "fmt"

//定义单链表

type ListNode struct {
	Val int
	Next *ListNode
}

//功能：创建单链表
func cre_link(data [5]int, node *ListNode) *ListNode {
	head := &ListNode{Val:data[0]}
	tail := head		// 需要头尾两个指针
	for i:= 1; i < len(&data); i++{
		// 方法一 数组直接构建链表
		tail.Next = &ListNode{Val:data[i]}
		tail = tail.Next
		// 方法二：追加构建链表
		//head.Append(list[i])

	}
	return head

}

// 追加构建链表
func (h *ListNode)Append(i int)  {
	for h.Next != nil{
		h = h.Next
	}
}

// 遍历链表
func (h *ListNode)Show()  {
	fmt.Println(h.Val)
	for h.Next != nil{
		h = h.Next
		fmt.Println(h.Val)
	}
}

//功能：删除单链表中的重复项
func deleteDuplicates(head *ListNode) *ListNode {
	node := head						// 记住链表头的位置
	for node != nil && node.Next != nil{
		if node.Val == node.Next.Val{	//遇到相同的两项，删掉
			node.Next = node.Next.Next
		}else{							// 遇到不同的两项，过
			node = node.Next
		}
	}
	return head
}

func main()  {
	var node *ListNode
	data := [5]int{1, 1, 2, 3, 3}
	node = cre_link(data,node)
	
	node.Show()
	node  = deleteDuplicates(node)
	node.Show()

}