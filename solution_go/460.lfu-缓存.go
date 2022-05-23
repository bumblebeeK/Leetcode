package solution

import "container/list"

/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

// @lc code=start
type LFUCache struct {
	capacity, min int
	mapKvs        map[int]element
	mapFreq       map[int]*list.List
}

type element struct {
	value int
	count int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		mapKvs:   make(map[int]element, capacity),
		mapFreq:  make(map[int]*list.List, capacity),
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	ele, ok := this.mapKvs[key]
	if !ok {
		return -1
	}
	this.mapKvs[key] = element{value: ele.value, count: ele.count + 1}
	cnt := ele.count

	for e := this.mapFreq[cnt].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.mapFreq[cnt].Remove(e)
			break
		}
	}
	if this.mapFreq[cnt+1] == nil {
		this.mapFreq[cnt+1] = list.New()
	}
	this.mapFreq[cnt+1].PushBack(key)

	if cnt == this.min && this.mapFreq[cnt].Len() == 0 {
		this.min++
	}
	return ele.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	_, ok := this.mapKvs[key]
	if ok {
		this.mapKvs[key] = element{value: value, count: this.mapKvs[key].count}
		this.Get(key)
		return
	}
	if len(this.mapKvs) >= this.capacity {
		l := this.mapFreq[this.min]
		delete(this.mapKvs, l.Front().Value.(int))
		l.Remove(l.Front())
	}
	this.mapKvs[key] = element{value: value, count: 1}
	if this.min != 1 {
		this.mapFreq[1] = list.New()
	}
	this.mapFreq[1].PushBack(key)
	this.min = 1
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
