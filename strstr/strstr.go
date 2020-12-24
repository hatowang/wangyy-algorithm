package main

import "fmt"

/*给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//想法如果needle等于0，则返回0，
//滑动窗口，比较切片的值
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i,j int
	for i = 0; i < len(haystack) - len(needle) + 1; i++ {
		for j=0; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}


func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Print(strStr(haystack, needle))
}