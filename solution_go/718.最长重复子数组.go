package solution

/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	res, dp := 0, make([][]int, len(nums1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}

	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res
}

// @lc code=end
