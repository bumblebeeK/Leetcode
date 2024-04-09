package solutiongo

import (
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 */

// @lc code=start
type RandomizedSet struct {
	set map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: map[int]int{},
		arr: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, has := this.set[val]; has {
		return false
	}
	this.set[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx := -1
	if i, has := this.set[val]; !has {
		return false
	} else {
		idx = i
	}
	last := len(this.arr) - 1
	this.arr[idx] = this.arr[last]
	this.set[this.arr[idx]] = idx
	this.arr = this.arr[:last]
	delete(this.set, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
