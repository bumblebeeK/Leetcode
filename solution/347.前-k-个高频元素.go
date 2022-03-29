package solution

import "container/heap"

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}

	q := &minStack{}
	heap.Init(q)
	for v, cnt := range m {
		heap.Push(q, element{value: v, count: cnt})
		if q.Len() > k {
			heap.Pop(q)
		}
	}

	ans := make([]int, k, k)
	for i := 0; i < k; i++ {
		v := heap.Pop(q)
		if s, ok := v.(element); ok {
			ans[i] = s.value
		}
	}
	return ans
}

type element struct {
	value int
	count int
}

type minStack []element

func (m minStack) Len() int { return len(m) }

func (m minStack) Less(i, j int) bool { return m[i].count < m[j].count }

func (m minStack) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *minStack) Pop() interface{} {
	old := *m
	ln := len(old)
	x := old[ln-1]
	*m = old[:ln-1]
	return x
}

func (m *minStack) Push(x interface{}) {
	*m = append(*m, x.(element))
}

// @lc code=end
