package solution

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 */

// @lc code=start
func readBinaryWatch(turnedOn int) []string {
	ans := []string{}
	arr := []int{1, 2, 4, 8, 1, 2, 4, 8, 16, 32}
	var dfs func(h, m, cnt, idx int)
	dfs = func(h, m, cnt, idx int) {
		if cnt == 0 {
			var minute string
			hour := strconv.Itoa(h)
			if m < 10 {
				minute = fmt.Sprintf("0%d", m)
			} else {
				minute = strconv.Itoa(m)
			}
			ans = append(ans, hour+":"+minute)
		}
		for i := idx; i < 10 && cnt > 0; i++ {
			if i < 4 && h+arr[i] < 12 {
				dfs(h+arr[i], m, cnt-1, i+1)
			}
			if i > 3 && m+arr[i] < 60 {
				dfs(h, arr[i]+m, cnt-1, i+1)
			}
		}
	}
	dfs(0, 0, turnedOn, 0)
	return ans
}

// @lc code=end
