package solution

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m + n - 1
	m--
	n--
	for n >= 0 {
		for m >= 0 && nums1[m] > nums2[n] {
			nums1[l] = nums1[m]
			l--
			m--
		}
		nums1[l] = nums2[n]
		n--
		l--
	}

}

// @lc code=end
