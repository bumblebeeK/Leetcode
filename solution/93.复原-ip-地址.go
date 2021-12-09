package solution

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start

func restoreIpAddresses(s string) []string {
	res := []string{}
	var dfs = func(subArr []string, start int) {}
	dfs = func(subArr []string, start int) {
		if len(subArr) == 4 {
			if len(s) == start {
				res = append(res, strings.Join(subArr, "."))
				return
			} else {
				return
			}
		}
		for i := 1; i <= 3; i++ {
			if start+i-1 >= len(s) {
				return
			}
			if i != 1 && s[start] == 48 {
				return
			}
			str := s[start : start+i]
			if ip, _ := strconv.Atoi(str); ip > 0xff {
				return
			}
			subArr = append(subArr, str)
			dfs(subArr, start+i)
			subArr = subArr[:len(subArr)-1]
		}
	}
	dfs([]string{}, 0)
	return res
}

// @lc code=end
