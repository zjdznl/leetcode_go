package leetcode

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	graph := make(map[int][]int)
	for i, m := range manager {
		graph[m] = append(graph[m], i)
	}

	return dfs(graph, headID, informTime)
}

func dfs(graph map[int][]int, start int, informTime []int) int {
	maxTime := 0
	if len(graph[start]) == 0 {
		return informTime[start]
	}
	for _, v := range graph[start] {
		maxTime = max(maxTime, dfs(graph, v, informTime))
	}
	return maxTime + informTime[start]
}

//
//func max(i1, i2 int) int {
//	if i1 > i2 {
//		return i1
//	}
//	return i2
//}
