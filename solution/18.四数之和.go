package solution

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}

	counter := map[int]int{}
	for _, v := range nums {
		counter[v]++
	}
	uniqArr := []int{}
	for k := range counter {
		uniqArr = append(uniqArr, k)
	}

	for i := 0; i < len(uniqArr); i++ {
		if uniqArr[i] == target/4 && counter[uniqArr[i]] >= 4 {
			res = append(res,[]int{uniqArr[i],uniqArr[i],uniqArr[i],uniqArr[i])
		}
		for j := i+1; j < count; j++ {
			if uniqArr[i]*3 +uniqArr[j] == target {
				
			}
		}
	}

}

// @lc code=end
