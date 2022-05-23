package solution

import "container/list"

/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N 叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	var res []int
	if root != nil {
		stack := list.New()
		stack.PushFront(root)
		for stack.Len() > 0 {
			curr := stack.Remove(stack.Front()).(*Node)
			if curr != nil {
				res = append(res, curr.Val)
			}
			if curr.Children != nil {
				for _, v := range curr.Children {
					stack.PushFront(v)
				}
			}
		}
		reverse(res)
	}
	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// @lc code=end
type Node struct {
	Val      int
	Children []*Node
}
