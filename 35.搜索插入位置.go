/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := left + (right-left)>>1
		if target <= nums[mid] {
			ans = mid
			right = mid - 1

		} else {
			left = mid + 1
		}

	}
	return ans
}

// @lc code=end

