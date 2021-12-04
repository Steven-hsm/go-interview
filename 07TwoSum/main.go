package main

func main() {

}
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if target-v == nums[k] {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for index, num := range nums {
		sencondNum := target - num
		if _, ok := numMap[sencondNum]; ok {
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
