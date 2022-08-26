package solution

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start

var dic = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var res = []string{}

	if digits == "" {
		return res
	}

	var combine func(pre, str string)
	combine = func(pre, str string) {
		if len(str) == 0 {
			res = append(res, pre)
			return
		}
		arr := dic[str[0]]
		for _, v := range arr {
			combine(pre+v, str[1:])
		}
	}
	combine("", digits)
	return res
}

// @lc code=end
