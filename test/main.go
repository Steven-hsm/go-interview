package main

import "fmt"

/**
给定一个整数数组 `nums` 和一个目标值 `target`，请你在该数组中找出和为目标值的那 **两个** 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
*/
func main() {
	fmt.Print(twoSum([]int{1, 1, 3, 3, 3}, 2))
}

func twoSum(nums []int, target int) []int {
	//Map key=value value= index
	numMap := make(map[int]int, len(nums))
	for index, num := range nums {
		//遍历所有的数据，第二个数就是target - num
		sencondNum := target - num
		//如果第二个数已经放到Map中了，则可以取出第二个数
		if _, ok := numMap[sencondNum]; ok {
			//第二个数是本身则需要继续遍历
			if index == numMap[sencondNum] {
				continue
			} else if index > numMap[sencondNum] {
				return []int{numMap[sencondNum], index}
			}
			return []int{index, numMap[sencondNum]}
		}
		numMap[num] = index
	}
	return []int{0, 0}
}
