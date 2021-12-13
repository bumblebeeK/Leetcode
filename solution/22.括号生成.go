package solution

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}

	if n == 0 {
		return res
	}
	var dfs = func(s string, l, r int) {}
	dfs = func(s string, l, r int) {
		if l == n && r == n {
			res = append(res, s)
			return
		}
		if l < n {
			dfs(s+"(", l+1, r)
		}
		if r < n && r < l {
			dfs(s+")", l, r+1)
		}
	}
	dfs("", 0, 0)
	return res
}

// @lc code=end
