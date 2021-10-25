/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	ptr := len(digits) - 1
	for ; ptr >= 0; ptr-- {
		if digits[ptr] < 9 {
			digits[ptr] += 1
			return digits
		} else {
			digits[ptr] = 0
		}
	}
	return append([]int{1}, digits...)

}

// @lc code=end

