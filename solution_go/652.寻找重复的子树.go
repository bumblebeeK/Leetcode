package solution

/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	repeat := map[*TreeNode]struct{}{}
	type pair struct {
		node *TreeNode
		idx  int
	}
	idx := 0
	seen := map[[3]int]pair{}
	ans := []*TreeNode{}
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		p := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}

		if pair, ok := seen[p]; ok {
			repeat[pair.node] = struct{}{}
			return pair.idx
		}

		idx++
		seen[p] = pair{node, idx}
		return idx

	}

	dfs(root)
	for v := range repeat {
		ans = append(ans, v)
	}
	return ans
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
