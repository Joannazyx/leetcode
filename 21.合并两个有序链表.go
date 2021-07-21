/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := ListNode{}
	pre.Val = -100
	pre.Next = l1

	pt2 := l2
	cur := &pre
	pt1 := l1

	for pt1 != nil && pt2 != nil {
		if pt1.Val <= pt2.Val {
			cur.Next = pt1
			cur = cur.Next
			pt1 = pt1.Next

		} else {
			cur.Next = pt2
			cur = cur.Next
			pt2 = pt2.Next
		}
	}

	if pt1 != nil {
		cur.Next = pt1
	} else {
		cur.Next = pt2
	}

	return pre.Next

}

// @lc code=end

