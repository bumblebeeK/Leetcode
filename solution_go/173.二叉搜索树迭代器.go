package solution

/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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
type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {

	c := BSTIterator{
		stack: []*TreeNode{},
		cur:   root,
	}
	return c

}

func (this *BSTIterator) Next() int {
	if this.cur == nil {
		return -1
	}
	for this.cur.Left != nil {
		this.stack = append(this.stack, this.cur)
		left := this.cur.Left
		this.cur.Left = nil
		this.cur = left
	}
	ans := this.cur.Val

	if this.cur.Right != nil {
		this.cur = this.cur.Right
	} else if len(this.stack) != 0 {
		this.cur = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
	} else {
		this.cur = nil
	}
	return ans
}

func (this *BSTIterator) HasNext() bool {
	return this.cur != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
