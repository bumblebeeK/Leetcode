package solution

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	arr  []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	if len(this.mins) == 0 || this.mins[len(this.mins)-1] >= val {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.arr) == 0 {
		panic("no element found in this array.")
	}
	if this.arr[len(this.arr)-1] == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	if len(this.arr) == 0 {
		panic("no element found in this array.")
	}
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.arr) == 0 {
		panic("no element found in this array.")
	}
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
