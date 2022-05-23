package solution

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	length := len(nums)
	l, r, ans := 1, length-1, -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0	
		for _, v := range nums {
			if v <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}

// @lc code=end
