package solution

import "sort"

/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 */

// @lc code=start

var c, a []int

func countSmaller(nums []int) []int {
	counts := []int{}
	c = make([]int, len(nums)+1)
	discretization(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		id := getId(nums[i])
		counts = append(counts, query(id-1))
		update(id)
	}

	for i := 0; i < len(counts)/2; i++ {
		counts[i], counts[len(nums)-i-1] = counts[len(nums)-i-1], counts[i]
	}
	return counts
}

func query(pos int) int {
	ret := 0
	for pos > 0 {
		ret += c[pos]
		pos -= lowbit(pos)
	}
	return ret
}

func getId(x int) int {
	return sort.SearchInts(a, x) + 1
}

func update(x int) {
	for x < len(c) {
		c[x]++
		x += lowbit(x)
	}
}

func lowbit(x int) int {
	return x & -x
}

func discretization(nums []int) {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	a = make([]int, 0, len(nums))
	for v := range m {
		a = append(a, v)
	}
	sort.Ints(a)
}

// @lc code=end
