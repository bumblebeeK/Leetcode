package solution

/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start
func sortArray(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}
	minV, maxV := nums[0], nums[0]
	for _, v := range nums {
		if minV > v {
			minV = v
		}
		if maxV < v {
			maxV = v
		}
	}
	if minV == maxV {
		return nums
	}
	bucket := make([]int, maxV-minV+1)
	for _, num := range nums {
		bucket[num-minV]++
	}
	index := 0
	for i := 0; i < len(bucket); i++ {
		for j := 0; j < bucket[i]; j++ {
			nums[index] = i + minV
			index++
		}
	}
	return nums
}

// @lc code=end
