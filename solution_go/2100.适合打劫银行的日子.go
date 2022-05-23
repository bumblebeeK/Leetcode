package solution

/*
 * @lc app=leetcode.cn id=2100 lang=golang
 *
 * [2100] 适合打劫银行的日子
 */

// @lc code=start
func goodDaysToRobBank(security []int, time int) []int {
	var res []int
	ln := len(security)
	if time*2+1 > ln {
		return res
	}
	if time == 0 {
		for i := 0; i < ln; i++ {
			res = append(res, i)
		}
		return res
	}
	var x, y int
	for i := 1; i <= time; i++ {
		if security[i] > security[i-1] {
			x++
		}
	}
	for j := time + 1; j <= time*2; j++ {
		if security[j] < security[j-1] {
			y++
		}
	}

	for i, j, cursor := 0, time*2, time; ; cursor++ {
		if x == 0 && y == 0 {
			res = append(res, cursor)
		}
		if j+1 == ln {
			break
		}
		if security[i+1] > security[i] {
			x--
		}
		if security[j+1] < security[j] {
			y++
		}
		if security[cursor+1] > security[cursor] {
			x++
		}
		if security[cursor+1] < security[cursor] {
			y--
		}
		i++
		j++
	}
	return res
}

// @lc code=end
