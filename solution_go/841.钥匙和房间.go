package solution

/*
 * @lc app=leetcode.cn id=841 lang=golang
 *
 * [841] 钥匙和房间
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	b := make([]bool, n)
	b[0] = true
	stack := []int{0}
	vist := 1
	for len(stack) > 0 {
		l := len(stack)
		for l > 0 {
			cur := stack[0]
			stack = stack[1:]
			for _, v := range rooms[cur] {
				if !b[v] {
					stack = append(stack, v)
					b[v] = true
					vist++
				}
			}
			l--
		}
	}
	return vist == n
}

// @lc code=end
