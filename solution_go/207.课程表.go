package solution

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		valid  = true
		dfs    func(u int)
		edges  = make([][]int, numCourses)
		visted = make([]int, numCourses)
	)
	dfs = func(u int) {
		visted[u] = 1
		for _, v := range edges[u] {
			if visted[v] == 1 {
				valid = false
				return
			} else if visted[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			}
		}

		visted[u] = 2
	}

	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visted[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// @lc code=end
