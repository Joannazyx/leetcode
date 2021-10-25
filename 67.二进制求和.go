/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	res := ""
	carry := 0
	ptra := len(a)
	ptrb := len(b)
	n := Max(ptra, ptrb)
	for i := 0; i < n; i++ {
		if i < ptra && string(a[ptra-i-1]) == "1" {
			carry += 1
		}
		if i < ptrb && string(b[ptrb-i-1]) == "1" {
			carry += 1
		}
		res = strconv.Itoa(carry%2) + res
		carry /= 2

	}
	if carry > 0 {
		res = "1" + res
	}
	return res

}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end

