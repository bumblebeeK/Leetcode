package solution

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var traceBack func(path []int, start int)
	traceBack = func(path []int, start int) {
		arr := make([]int, len(path))
		copy(arr, path)
		ans = append(ans, arr)
		for i := start; i < n; i++ {
			path = append(path, nums[i])
			traceBack(path, i+1)
			path = path[:len(path)-1]
		}
	}

	traceBack([]int{}, 0)
	return ans
}

// @lc code=end
