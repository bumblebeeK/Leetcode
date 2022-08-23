package solution

import "strings"

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	stack := []string{}
	for _, v := range strings.Split(path, "/") {
		if v == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if v != "." && v != "" {
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}

// @lc code=end
