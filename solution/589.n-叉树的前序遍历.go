package solution

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	if root.Children != nil && len(root.Children) > 0 {
		for _, v := range root.Children {
			arr := preorder(v)
			res = append(res, arr...)
		}
	}

	return res
}

// @lc code=end
type Node struct {
	Val      int
	Children []*Node
}
