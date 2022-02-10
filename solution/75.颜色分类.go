package solution

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start

// Hash方法

// func sortColors(nums []int) {
// 	m := map[int]int{}
// 	for _, num := range nums {
// 		m[num]++
// 	}

// 	for i, j := 0, 0; i < 3; i++ {
// 		for m[i] != 0 {
// 			m[i]--
// 			nums[j] = i
// 			j++
// 		}
// 	}
// }

// 计数抢占法
func sortColors(nums []int) {
	zero, one := 0, 0
	for i, num := range nums {
		nums[i] = 2
		if num < 2 {
			nums[one] = 1
			one++
		}
		if num < 1 {
			nums[zero] = 0
			zero++
		}
	}
}

// @lc code=end
