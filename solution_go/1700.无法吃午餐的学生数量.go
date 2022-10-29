package solution

/*
 * @lc app=leetcode.cn id=1700 lang=golang
 *
 * [1700] 无法吃午餐的学生数量
 */

// @lc code=start
func countStudents(students []int, sandwiches []int) int {
	oneStu := 0
	oneSand := 0
	for _, v := range students {
		if v == 1 {
			oneStu++
		}
	}

	for _, v := range sandwiches {
		if v == 1 {
			oneSand++
		}
	}
	if oneSand == oneStu {
		return 0
	}
	if oneSand > oneStu {
		cnt := 0
		for i, v := range sandwiches {
			if v == 1 {
				cnt++
			}
			if cnt == oneStu+1 {
				return len(sandwiches) - i
			}
		}
	}
	cnt := 0
	for i, v := range sandwiches {
		if v == 0 {
			cnt++
		}
		if cnt == len(students)-oneStu+1 {
			return len(sandwiches) - i
		}
	}
	return 0
}

// @lc code=end
