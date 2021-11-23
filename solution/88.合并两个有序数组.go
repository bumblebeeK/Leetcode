package solution

import "fmt"

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	l := len(nums1) - 1
	fmt.Println(l)
	n--
	m--
	for n >= 0 {
		for m >= 0 && nums2[n] < nums1[m] {
			nums1[l] = nums1[m]
			nums1[m] = 0
			l--
			m--
		}
		nums1[l] = nums2[n]
		l--
		n--
	}
}

// @lc code=end
