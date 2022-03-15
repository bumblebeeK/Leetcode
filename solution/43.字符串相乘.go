package solution

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */
// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ln1, ln2 := len(num1), len(num2)
	ln := ln1 + ln2
	res := make([]byte, ln)
	for i := ln1 - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		for j := ln2 - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			sum := n1*n2 + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	if res[0] == 0 {
		res = res[1:]
	}
	for i, _ := range res {
		res[i] += '0'
	}
	return string(res)
}

// @lc code=end
