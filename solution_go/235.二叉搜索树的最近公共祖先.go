package solution

/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	if root == nil {
		return res
	}
	var lca func(root, p, q *TreeNode)
	lca = func(root, p, q *TreeNode) {
		if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
			res = root
		} else if root.Val > p.Val && root.Val > q.Val {
			lca(root.Left, p, q)
		} else {
			lca(root.Right, p, q)
		}
	}
	lca(root, p, q)
	return res
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
