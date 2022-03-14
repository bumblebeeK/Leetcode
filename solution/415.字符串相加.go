package solution

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	ln := max(len(num1), len(num2))
	byteDance := make([]byte, ln+1)
	var i, j = len(num1) - 1, len(num2) - 1
	for i >= 0 || j >= 0 {
		if i >= 0 && j >= 0 {
			sum := num1[i] + num2[j] - 2*'0' + byteDance[max(i, j)+1]
			byteDance[max(i, j)] = sum / 10
			byteDance[max(i, j)+1] = sum%10 + '0'
			i--
			j--
			continue
		}
		if i >= 0 {
			sum := num1[i] - '0' + byteDance[i+1]
			byteDance[i] = sum / 10
			byteDance[i+1] = sum%10 + '0'
			i--
			continue
		}
		if j >= 0 {
			sum := num2[j] - '0' + byteDance[j+1]
			byteDance[j] = sum / 10
			byteDance[j+1] = sum%10 + '0'
			j--
		}
	}
	if byteDance[0] == 0 {
		return string(byteDance[1:])
	}
	byteDance[0] += '0'
	return string(byteDance)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
