package solution

/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	return searchX(nums, 0, len(nums)-1, target)
}

func searchX(nums []int, start, end, target int) int {
	if end-start < 0 {
		return -1
	}
	middle := (end-start)/2 + start
	if nums[middle] == target {
		return middle
	} else if nums[middle] > target {
		return searchX(nums, start, middle-1, target)
	} else {
		return searchX(nums, middle+1, end, target)
	}

}

// @lc code=end
