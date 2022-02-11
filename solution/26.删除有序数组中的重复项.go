package solution

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start

// 懒指针
// func removeDuplicates(nums []int) int {
// 	l := len(nums)
// 	if l < 1 {
// 		return l
// 	}
// 	fast, slow := 1, 0
// 	for fast < l {
// 		if nums[fast] != nums[slow] {
// 			slow++
// 			nums[slow] = nums[fast]
// 		}
// 		fast++
// 	}
// 	return slow + 1
// }

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 1 {
		return l
	}
	fast, slow := 1, 1
	for fast < l {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow

}

// @lc code=end
