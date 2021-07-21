/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lptr, rptr := 0, 0
	for lptr < len(nums)-1 {
		for nums[lptr] == nums[rptr] {
			rptr++
			if rptr == len(nums) {
				return lptr + 1
			}
		}
		nums[lptr+1] = nums[rptr]
		lptr++

	}
	return lptr + 1
}

// @lc code=end

