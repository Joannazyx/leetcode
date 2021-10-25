/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	ptr := len(s) - 1
	length := 0
	for ; s[ptr] == ' '; ptr-- {

	}
	for ptr >= 0 && s[ptr] != ' ' {
		length++
		ptr--

	}
	return length
}

// @lc code=end

