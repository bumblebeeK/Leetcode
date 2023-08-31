package solution

/*
 * @lc app=leetcode.cn id=2540 lang=golang
 *
 * [2540] 最小公共值
 */

// @lc code=start
func getCommon(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	i, j := 0, 0
	for i < n1 && j < n2 {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			return nums1[i]
		}
	}
	return -1
}

// @lc code=end
