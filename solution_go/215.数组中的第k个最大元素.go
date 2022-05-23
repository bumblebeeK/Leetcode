package solution

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start

func findKthLargest(nums []int, k int) int {
	insertSort(nums, 0, len(nums)-1)
	return nums[k-1]
}

// func swapEle(nums []int, index1, index2 int) {
// 	tmp := nums[index1]
// 	nums[index1] = nums[index2]
// 	nums[index2] = tmp
// }

func insertSort(nums []int, left, right int) {
	for i := right - 1; i >= left; i-- {
		tmp := nums[i]
		j := i
		for j < right && nums[j+1] > tmp {
			nums[j] = nums[j+1]
			j++
		}
		nums[j] = tmp
	}
}

// func quickSort(nums []int, left, right int) {
// 	if right-left <= INSERTION_SORT_THRESHOLD {
// 		insertSort(nums, left, right)
// 		return
// 	}
// 	rand.Seed(time.Now().UnixNano())
// 	rIndex := rand.Intn(right-left+1) + left
// 	swapEle(nums, left, rIndex)
// 	piovt := nums[left]
// 	gt := left
// 	lt := right + 1
// 	i := left + 1
// 	for i < lt {
// 		if nums[i] == piovt {
// 			i++
// 		}
// 		if nums[i] < piovt {
// 			lt--
// 			swapEle(nums, i, lt)
// 		} else {
// 			gt++
// 			swapEle(nums, i, gt)
// 			i++
// 		}
// 	}
// 	swapEle(nums, left, gt)
// 	quickSort(nums, left, i-1)
// 	quickSort(nums, i+1, right)
// }

// @lc code=end

// if right-left <= INSERTION_SORT_THRESHOLD {
// 	insertSort(nums, left, right)
// 	return
// }
// rand.Seed(time.Now().UnixNano())
// rIndex := rand.Intn(right-left+1) + left
// swapEle(nums, left, rIndex)

// pivot := nums[left]
// lt := left
// gt := right + 1
// i := left + 1
// for i < gt {
// 	if nums[i] > pivot {
// 		gt--
// 		swapEle(nums, i, gt)
// 	} else if pivot == nums[i] {
// 		i++
// 	} else {
// 		lt++
// 		swapEle(nums, i, lt)
// 	}
// }
// swapEle(nums, left, lt)
// quickSort(nums, left, lt-1)
// quickSort(nums, gt, right)

func insertSortPractice(nums []int, left, right int) {
	for i := right - 1; i >= 0; i-- {
		tmp := nums[i]
		j := i
		for j < right && nums[j+1] > tmp {
			nums[j] = nums[j+1]
		}
		nums[j] = tmp
	}
}
