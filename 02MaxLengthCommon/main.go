package main

import (
	"fmt"
	"strings"
)

/**
*编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
 */
func main() {
	str := []string{"flow", "flower", "flat"}
	fmt.Print(longestCommonPrefix(str))

}

/**
 * 分析，注意：最长公共前缀
 * 取第一个元素为首次比较的前缀，先找到与第二个元素最长公共前缀，然后再依次匹配第三个、第四个...
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	//取第一个元素为第一次的最长公共前缀比较
	prefix := strs[0]
	for _, str := range strs {
		//遍历元素，依次往后找最长公共前缀
		for strings.Index(str, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			//如果不满足，prefix从最后减少一个字符串
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
