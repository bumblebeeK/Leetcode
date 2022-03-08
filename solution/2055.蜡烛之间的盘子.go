package solution

import "fmt"

/*
 * @lc app=leetcode.cn id=2055 lang=golang
 *
 * [2055] 蜡烛之间的盘子
 */

// @lc code=start
func platesBetweenCandles(s string, queries [][]int) []int {
	var candle []int
	ans := make([]int, len(queries))
	sum := make([]int, len(s)+1)
	for i, v := range s {
		if v == '|' {
			candle = append(candle, i)
			sum[i+1] = sum[i]
		} else {
			sum[i+1] = sum[i] + 1
		}

	}
	if len(candle) == 0 {
		return ans
	}
	for i, v := range queries {
		a, b, ll, rr := v[0], v[1], -1, -1
		l, r := 0, len(candle)-1
		for l < r {
			middle := (l + r) >> 1
			if candle[middle] >= a {
				r = middle
			} else {
				l = middle + 1
			}
		}
		if candle[r] >= a {
			ll = candle[r]
		} else {
			continue
		}
		l, r = 0, len(candle)-1
		for l < r {
			middle := (l + r + 1) >> 1
			if candle[middle] <= b {
				l = middle
			} else {
				r = middle - 1
			}
		}
		if candle[r] <= b {
			rr = candle[r]
		} else {
			continue
		}
		if ll <= rr {
			ans[i] = sum[rr] - sum[ll]
		}
	}
	return ans
}

// @lc code=end
