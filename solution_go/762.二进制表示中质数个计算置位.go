package solution

/*
 * @lc app=leetcode.cn id=762 lang=golang
 *
 * [762] 二进制表示中质数个计算置位
 */

// @lc code=start
func countPrimeSetBits(left int, right int) int {
	ans := 0
	var dic = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	hash := make(map[int]bool, len(dic))
	for _, v := range dic {
		hash[v] = true
	}

	for i := left; i <= right; i++ {
		cnt, cur := 0, i
		for cur > 0 {
			cur -= (cur & -cur)
			cnt++
		}
		if hash[cnt] {
			ans++
		}
	}
	return ans
}

// @lc code=end
