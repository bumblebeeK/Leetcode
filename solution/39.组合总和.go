package solution

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	stack := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), stack...))
			return
		}
		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			stack = append(stack, candidates[idx])
			dfs(target-candidates[idx], idx)
			stack = stack[:len(stack)-1]
		}
	}
	dfs(target, 0)
	return ans
}

// @lc code=end
