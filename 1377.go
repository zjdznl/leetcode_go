package leetcode
//
//func frogPosition(n int, edges [][]int, t int, target int) float64 {
//	graph := make(map[int][]int)
//	for _, v := range edges {
//		graph[v[0]] = append(graph[v[0]], v[1])
//		graph[v[1]] = append(graph[v[1]], v[0])
//	}
//
//	return dfs(graph, 1, t, target)
//}
//
//func dfs(graph map[int][]int, start int, t int, target int) float64 {
//	if len(graph[t]) != 0 {
//		return 0
//	}
//}
//
//func inSlice(target int, slice []int) bool {
//	for _, v := range slice {
//		if v == target {
//			return true
//		}
//	}
//	return false
//}
