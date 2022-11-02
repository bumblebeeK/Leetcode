package solution

import "math"

/*
 * @lc app=leetcode.cn id=1620 lang=golang
 *
 * [1620] 网络信号最好的坐标
 */

// @lc code=start
func bestCoordinate(towers [][]int, radius int) []int {
	var ans []int
	x, y, val := 0, 0, 0
	for _, v := range towers {
		a, b, q := v[0]-radius, v[1]-radius, v[2]
		if a < 0 {
			a = 0
		}
		if b < 0 {
			b = 0
		}
		for i := a; i <= a+2*radius; i++ {
			for j := b; j <= b+2*radius; j++ {
				math.Sqrt()
			}
		}
	}

	return ans
}

// @lc code=end
