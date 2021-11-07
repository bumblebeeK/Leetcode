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
		if uniqArr[i] == 0 && counter[uniqArr[i]] >= 3 {
			res = append(res, []int{uniqArr[i], uniqArr[i], uniqArr[i]})
		}
		for j := i + 1; j < len(uniqArr); j++ {
			if uniqArr[i]*2+uniqArr[j] == 0 && counter[uniqArr[i]] > 1 {
				res = append(res, []int{uniqArr[i], uniqArr[i], uniqArr[j]})
			}
			if uniqArr[j]*2+uniqArr[i] == 0 && counter[uniqArr[j]] > 1 {
				res = append(res, []int{uniqArr[j], uniqArr[j], uniqArr[i]})
			}
			c := 0 - uniqArr[i] - uniqArr[j]
			if counter[c] > 0 && c > uniqArr[j] {
				res = append(res, []int{uniqArr[i], uniqArr[j], c})
			}
		}

	}
	return res
}

// @lc code=end
