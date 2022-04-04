package solution

import "math"

/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 */

// @lc code=start
type NumArray struct {
	vals []int
	size int
	sums []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	size := int(math.Sqrt(float64(n)))
	sums := make([]int, (size+n-1)/size)
	for i, v := range nums {
		sums[i/size] += v
	}
	return NumArray{size: size, vals: nums, sums: sums}
}

func (this *NumArray) Update(index int, val int) {
	this.sums[index/this.size] += val - this.vals[index]
	this.vals[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	ans := 0
	l1, l2 := left/this.size, right/this.size
	if l1 == l2 {
		for i := left; i <= right; i++ {
			ans += this.vals[i]
		}
		return ans
	}
	for i := l1 + 1; i < l2; i++ {
		ans += this.sums[i]
	}
	for i := left; i < (l1+1)*this.size; i++ {
		ans += this.vals[i]
	}
	for i := l2 * this.size; i <= right; i++ {
		ans += this.vals[i]
	}
	return ans
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end
