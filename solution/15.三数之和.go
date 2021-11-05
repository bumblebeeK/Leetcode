package solution

import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	counter := map[int]int{}
	for _, v := range nums {
		counter[v]++

	}
	uniqArr := []int{}
	for k := range counter {
		uniqArr = append(uniqArr, k)
	}

	sort.Ints(uniqArr)
	for i := 0; i < len(uniqArr); i++ {
		if uniqArr[i] == 0 && counter[uniqArr[i]] > 3 {
			res = append(res, []int{})
		}
	}
	return res
}

// @lc code=end
