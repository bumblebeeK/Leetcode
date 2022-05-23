package solution

/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
var size int64

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := map[int64]int64{}
	size = int64(t) + 1
	for i, v := range nums {
		u := v
		idx := getIndex(int64(u))
		if _, exist := m[idx]; exist {
			return true
		}
		l, r := idx-1, idx+1
		if _, exist := m[l]; exist && int64(u)-m[l] <= int64(t) {
			return true
		}
		if _, exist := m[r]; exist && m[r]-int64(u) <= int64(t) {
			return true
		}
		m[idx] = int64(u)
		if i >= k {
			delete(m, getIndex(int64(nums[i-k])))
		}
	}
	return false
}

func getIndex(u int64) int64 {
	if u >= 0 {
		return u / size
	} else {
		return (u+1)/size - 1
	}
}

// @lc code=end
