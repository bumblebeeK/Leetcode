package solution

/*
 * @lc app=leetcode.cn id=1773 lang=golang
 *
 * [1773] 统计匹配检索规则的物品数量
 */

// @lc code=start
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	m := map[byte]int{'t': 0, 'c': 1, 'n': 2}[ruleKey[0]]
	ans := 0
	for _, v := range items {
		if v[m] == ruleValue {
			ans++
		}
	}
	return ans
}

// @lc code=end
