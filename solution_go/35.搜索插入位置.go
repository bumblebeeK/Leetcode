package solution

import "fmt"

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := end - start>>1 + start
		if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	fmt.Println(start)
	fmt.Println(end)
	return start
}

// @lc code=end
