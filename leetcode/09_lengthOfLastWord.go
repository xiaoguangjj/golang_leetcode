/*
58. 最后一个单词的长度
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。



示例:

输入: "Hello World"
输出: 5

*/

package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	str := strings.Split(s," ")
	fmt.Printf("str: %s, len: %d, cut-1: %s\n", str, len(str)-1,str[len(str)-1])
	if str != nil {
		fmt.Printf("29\n")
		return len(str[len(str)-1])
	}else {
		fmt.Printf("32\n")
		return 0
	}
}

func main(){
	//str := "Hello World"
	str := "a "
	fmt.Printf("长度为：%d\n",lengthOfLastWord(str))
}