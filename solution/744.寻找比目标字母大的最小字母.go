package solution

/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] 寻找比目标字母大的最小字母
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)
	if letters[right-1] <= target {
		return letters[left]
	}
	for left < right {
		mid := (left + right) >> 1
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return letters[left]
}

// @lc code=end

// 0 2 3 4 5 6 7   || first 0 last 7 mid 3
// 0 1 2 3 4 5     || first 0 last 6 mid 3
