package main

import "fmt"

/**
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
*/
func main() {
	fmt.Print(plusOne([]int{8, 9, 8, 9, 9, 9}))
}
func plusOne(digits []int) []int {
	//从最后一个元素遍历，逢十进一
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				result := []int{1}
				return append(result, digits...)
			}
		} else {
			digits[i] = digits[i] + 1
			break
		}
	}
	return digits
}

func plusOne2(digits []int) []int {
	var result []int
	addon := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += addon
		addon = 0
		if i == len(digits)-1 {
			digits[i]++
		}
		if digits[i] == 10 {
			addon = 1
			digits[i] = digits[i] % 10
		}
	}
	if addon == 1 {
		result = make([]int, 1)
		result[0] = 1
		result = append(result, digits...)
	} else {
		result = digits
	}
	return result
}
