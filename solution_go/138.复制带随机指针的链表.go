package solution

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

var cache map[*Node]*Node

func copyRandomList(head *Node) *Node {
	cache = map[*Node]*Node{}
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cache[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cache[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)

	return newNode
}

// @lc code=end
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
