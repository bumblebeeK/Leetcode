package solution

/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {

	if root == nil {
		return 0
	}
	maxDe := 1
	var dfs func(*Node, int) int
	dfs = func(n *Node, cur int) int {
		if n == nil {
			return 0
		}
		cur += 1
		if len(n.Children) == 0 {
			return cur
		}
		for _, v := range n.Children {
			maxDe = max(dfs(v, cur), maxDe)
		}

		return cur
	}
	dfs(root, 0)
	return maxDe

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

type Node struct {
	Val      int
	Children []*Node
}
