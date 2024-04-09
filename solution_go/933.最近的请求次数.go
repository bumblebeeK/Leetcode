package solutiongo

/*
 * @lc app=leetcode.cn id=933 lang=golang
 *
 * [933] 最近的请求次数
 */

// @lc code=start
type RecentCounter struct {
	stack []int
}

func Constructor() RecentCounter {
	return RecentCounter{stack: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	cnt := 0
	for _, v := range this.stack {
		if t-v <= 3000 {
			break
		}
		cnt++
	}
	this.stack = this.stack[cnt:]
	this.stack = append(this.stack, t)
	return len(this.stack)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end
