package solution

/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	stack := []*Node{root}
	for len(stack) > 0 {
		l := len(stack)
		curLevel := []int{}
		for l > 0 {
			cur := stack[0]
			stack = stack[1:]
			for _, v := range cur.Children {
				stack = append(stack, v)
			}
			curLevel = append(curLevel, cur.Val)
			l--
		}
		ans = append(ans, curLevel)
	}
	return ans

}

// @lc code=end

type Node struct {
	Val      int
	Children []*Node
}
