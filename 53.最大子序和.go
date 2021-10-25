/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	//greedy algorithm
	maxsum := nums[0]
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		if maxsum < tmp {
			maxsum = tmp

		}
		if tmp < 0 {
			tmp = 0

		}
	}
	return maxsum

}

// @lc code=end

