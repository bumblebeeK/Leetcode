package solution

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	return searchX(nums, 0, len(nums)-1, target)
}
func searchX(nums []int, start, end, target int) []int {

	if start > end {
		return []int{-1, -1}
	}
	middle := (end-start)/2 + start
	if nums[middle] == target {
		o1, o2 := middle, middle
		for i := middle + 1; i <= end; i++ {
			if nums[i] != target {
				break
			}
			if nums[i] == target {
				o2 = i
			}
		}
		for i := middle - 1; i >= 0; i-- {
			if nums[i] != target {
				break
			}
			if nums[i] == target {
				o1 = i
			}
		}
		return []int{o1, o2}
	} else if nums[middle] > target {
		return searchX(nums, start, middle-1, target)
	} else {
		return searchX(nums, middle+1, end, target)
	}
}

// @lc code=end
