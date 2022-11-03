package solution

import "math"

/*
 * @lc app=leetcode.cn id=1620 lang=golang
 *
 * [1620] 网络信号最好的坐标
 */

// @lc code=start
func bestCoordinate(towers [][]int, radius int) []int {
	x, y, val := 0, 0, 0
	grid := make([][]int, 101)
	for i := range grid {
		grid[i] = make([]int, 101)
	}
	for _, v := range towers {
		a, b, q := v[0], v[1], v[2]
		for i := int(math.Max(0, float64(a-radius))); i <= a+radius; i++ {
			for j := int(math.Max(0, float64(b-radius))); j <= b+radius; j++ {
				d := math.Sqrt(float64((i-a)*(i-a) + (j-b)*(j-b)))
				if d > float64(radius) {
					continue
				}
				grid[i][j] += int(math.Floor(float64(float64(q) / (1 + d))))
				if grid[i][j] > val {
					x, y, val = i, j, grid[i][j]
				} else if grid[i][j] == val {
					if i < x || (i == x && j < y) {
						x, y = i, j
					}
				}
			}
		}
	}

	return []int{x, y}
}

// @lc code=end
