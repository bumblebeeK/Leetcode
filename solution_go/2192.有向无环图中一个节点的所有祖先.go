package solutiongo

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=2192 lang=golang
 *
 * [2192] 有向无环图中一个节点的所有祖先
 */

// @lc code=start
func getAncestors(n int, edges [][]int) [][]int {
	adjacency := make([][]int, n)
	anc := make([]map[int]bool, n)
	indeg := make([]int, n)
	for i := 0; i < n; i++ {
		anc[i] = map[int]bool{}
	}
	for _, edge := range edges {
		adjacency[edge[0]] = append(adjacency[edge[0]], edge[1])
		indeg[edge[1]]++
	}

	queue := []int{}
	for i, v := range indeg {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	ans := make([][]int, n)
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		for _, v := range adjacency[c] {
			for j := range anc[c] {
				anc[v][j] = true
			}
			anc[v][c] = true
			indeg[v]--
			if indeg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	for i := 0; i < n; i++ {
		for v := range anc[i] {
			ans[i] = append(ans[i], v)
		}
		sort.Ints(ans[i])
	}
	return ans
}

// @lc code=end
