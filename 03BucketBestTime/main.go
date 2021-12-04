package main

import "fmt"

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润
*/
func main() {
	//fmt.Print(buy([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Print(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Print(maxProfit([]int{1, 2, 3, 4, 5, 6}))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buyFlag := false
	buyMoney := 0
	totalMoney := 0
	for index, num := range prices {
		//如果没有买进，那么等待买进
		if !buyFlag {
			//下次要涨的话就买进
			if (index < len(prices)-1) && num < prices[index+1] {
				buyFlag = true
				buyMoney = num
			}
		} else {
			//已经买进，最后一笔直接卖出 或者下次要跌都要卖出
			if index == len(prices)-1 || num > prices[index+1] {
				totalMoney = num - buyMoney + totalMoney
				buyFlag = false
			}
		}
	}
	return totalMoney
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		//
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
