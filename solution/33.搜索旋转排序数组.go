package solution

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	res := -1
	faultage := -1

	var BinarySearch func(nums []int, i, j, target int)
	BinarySearch = func(nums []int, i, j, target int) {
		if i <= j {
			mid := (i + j) >> 1
			if nums[mid] == target {
				res = mid
				return
			}
			if nums[mid] > target {
				BinarySearch(nums, i, mid-1, target)
				return
			}
			BinarySearch(nums, mid+1, j, target)
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			faultage = i
		}
		if nums[i] == target {
			return i
		}
		if nums[i+1] == target {
			return i + 1
		}
	}
	if faultage != -1 {
		if nums[faultage] >= target && nums[0] <= target {
			BinarySearch(nums, 0, faultage, target)
		} else {
			BinarySearch(nums, faultage+1, len(nums)-1, target)
		}
	}
	return res
}

// @lc code=end
