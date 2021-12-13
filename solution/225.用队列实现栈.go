package solution

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	arr []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.arr = append(this.arr, x)
}

func (this *MyStack) Pop() int {
	l := len(this.arr)
	e := this.arr[l-1]
	this.arr = this.arr[:l-1]
	return e
}

func (this *MyStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MyStack) Empty() bool {
	if len(this.arr) > 0 {
		return false
	}
	return true
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
