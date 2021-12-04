package main

/*给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
要求使用空间复杂度为 O(1) 的 原地 算法。
*/
func main() {
}

/**
分析：一个数组反转两次就会还原
第一次反转改变位置，第二次反转就用来将数据还原
*/

func rotate(nums []int, k int) {
	//反转，以k为界的两段数组位置保证
	reverse(nums)
	//保证前k个元素还原
	reverse(nums[:k%len(nums)])
	//保证k个元素后面的数据还原
	reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
