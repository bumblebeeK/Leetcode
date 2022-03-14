package solution

/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	set := map[string]int{}
	for i, v := range list1 {
		set[v] = i
	}
	res := []string{}
	var min int
	for i2, v := range list2 {
		if i1, ok := set[v]; ok {
			if len(res) == 0 {
				res = append(res, v)
				min = i2 + i1
				continue
			} else if i1+i2 < min {
				min = i1 + i2
				res = []string{v}
			} else if i1+i2 == min {
				res = append(res, v)
			}
		}
	}
	return res
}

// @lc code=end
