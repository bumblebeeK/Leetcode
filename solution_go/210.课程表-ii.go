package solution

import "fmt"

/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		visted = make([]int, numCourses)
		edges  = make([][]int, numCourses)
		valid  = true
		result = []int{}
		dfs    func(u int)
	)

	dfs = func(u int) {
		visted[u] = 1
		for _, v := range edges[u] {
			if visted[v] == 1 {
				valid = false
				return
			} else if visted[v] == 0 {
				fmt.Println(v)
				dfs(v)
				if !valid {
					return
				}

			}
		}

		visted[u] = 2
		result = append(result, u)
	}

	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visted[i] == 0 {
			dfs(i)
			fmt.Println(i)
		}
	}

	if !valid {
		return []int{}
	}
	for i := 0; i < numCourses>>1; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}

// @lc code=end
