package solution

import "fmt"

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	set := NewSet()
	return dfs(n, set)
}

func dfs(n int, set *Set) bool {
	if n == 1 {
		return true
	}

	sum := 0
	for n > 0 {
		v := n % 10
		sum += v * v
		n = n / 10
	}
	if set.Has(sum) {
		return false
	}
	set.Add(sum)
	return dfs(sum, set)
}

type Set struct {
	items map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[int]struct{}),
	}
}

func (s *Set) Has(n int) bool {
	_, exists := s.items[n]
	return exists
}

func (s *Set) Add(n int) {
	s.items[n] = struct{}{}
}

// @lc code=end
