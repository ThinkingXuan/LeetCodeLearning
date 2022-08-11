package main

import "fmt"

//
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
//如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
//分析：小偷获取的是收入 num[i],  dp(i)前i个房屋能偷窃到的最高金额
//   当有一间房子是：只能偷窃这个房子，可以偷窃到最高金额   dp[0] = nums[0]
//   如果有两间房子，则 偷窃金额最大的那间房子，可以偷窃到最高金额  dp[1] = max(nums[0], nums[1])
//   如果k间房子，k>2, 有两种选择
//            1.  偷窃第k间房子，则不能偷窃第k-1间房子，所有偷窃总金额为 dp[k] = nums[k]+dp[k-2]
//            2.  不偷窃第k间房子，则偷窃最高金额为  dp[k] = dp[k-1]
//   所以动态转移方程为：
//            dp[i] = max{nums[i]+dp[i-2], dp[i-1]}
//
//    边界条件： dp[0] = nums[0]
//              dp[1] = max(nums[0], nums[1])
//
func main() {
	fmt.Println(rob([]int{2,7,9,3,1}))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int,len(nums))

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i:=2;i<n; i++ {
		dp[i] = max(dp[i-2] + nums[i], dp[i-1])
	}
	return dp[n-1]
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}