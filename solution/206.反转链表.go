package solution

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//当前节点在下一次循环时会变成下一个节点的Next，换句话说左节点在下一次循环时会变成右节点
//举个栗子 tmp只是用来记录下一次循环的点 重点研究head与pre
//n1为当前节点 第一次循环n1.Next = nil;pre = n1;head = n2    n1与主链脱钩 pre记录n1 n2->n3->n4->n5->n6
//n2为当前节点 第二次循环n2.Next = n1;pre = n2;head = n3     n2与主链脱钩 pre记录n2 n2->n1 n3->n4->n5->n6
//n3为当前节点 第三次循环n3.Next = n2;pre = n3;head = n4     n3与主链脱钩 pre记录n3 n3->n2->n1 n4->n5->n6
//总而言之 每次循环将当前点与原链脱钩接上前节点

func reverseList(head *ListNode) *ListNode {
	var tmp, pre *ListNode //
	for head != nil {
		tmp = head.Next //tmp记录下一次循环的点

		head.Next = pre //接上上一次循环的节点
		pre = head      //记录下当前点，以供给下一次循环接上

		//最后一次循环时，终节点赋给cur节点，完成倒置,终止循环

		head = tmp
	}
	return pre
}

// @lc code=end
