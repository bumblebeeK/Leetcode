package solution

/*
 * @lc app=leetcode.cn id=457 lang=golang
 *
 * [457] 环形数组是否存在循环
 */

// @lc code=start
func circularArrayLoop(nums []int) bool {
	n := len(nums)
	var next = func(idx int) int {
		return (((idx + nums[idx]) % n) + n) % n
	}

	for i, v := range nums {
		if v == 0 {
			continue
		}
		slow, fast := i, next(i)
		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[next(fast)] > 0 {
			if slow == fast {
				if slow != next(slow) {
					return true
				}
				break
			}
			slow, fast = next(slow), next(next(fast))
		}
		x := i
		for nums[x] != 0 {
			nums[x] = 0
			x = next(x)
		}

	}
	return false

}

// @lc code=end
