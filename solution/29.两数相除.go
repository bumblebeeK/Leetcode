package solution

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if (dividend == -2<<30) && (divisor == -1) {
		return 2<<30 - 1
	}
	absDevident := abs(dividend)
	absDevisor := abs(divisor)
	res := 0
	for absDevident-absDevisor >= 0 {
		expo := 0
		for absDevident-absDevisor<<expo >= 0 {
			expo++
		}
		res += 1 << (expo - 1)
		absDevident -= absDevisor << (expo - 1)

	}
	if (dividend > 0) == (divisor >= 0) {
		return res
	}
	return -res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// @lc code=end
