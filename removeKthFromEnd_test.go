package test

import (
	"testing"
)

const N = 2

func TestEmptyListNode(t *testing.T) {

	// tips: the length of test linkedList must greater than 2
	testCaseOne := []int{1, 2}
	testCaseOneToList := buildLinkedListFromArr(testCaseOne)
	expectation := []int{2}
	ans := buildArrFromList(removeNthFromEnd(testCaseOneToList, N))
	assertEqual(ans, expectation)

	testCaseTwo := []int{1, 2, 3, 4, 5}
	testCaseTwoToList := buildLinkedListFromArr(testCaseTwo)
	expectation = []int{1, 2, 3, 5}
	ans = buildArrFromList(removeNthFromEnd(testCaseTwoToList, N))
	assertEqual(ans, expectation)

}

func assertEqual(a, b []int) {
	if len(a) != len(b) {
		panic("a not equals b!")
	}
	for i, _ := range a {
		if a[i] != b[i] {
			panic("a not equals b!")
		}
	}
}

// buildLinkedListFromArr build *ListNode from int array
func buildLinkedListFromArr(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	return &ListNode{
		Val:  arr[0],
		Next: buildLinkedListFromArr(arr[1:]),
	}
}

// buildArrFromList build int arr from *ListNode
func buildArrFromList(node *ListNode) []int {
	arr := []int{}
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}
	return arr
}

// ListNode defined the ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd removes the Nth ListNode
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{}
	res.Next = head
	fast, slow := res, res

	// let the fast pointer move n steps first
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for nil != fast.Next {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return res.Next
}
