package leetcode

import "math"


func luckyNumbers(matrix [][]int) []int {
	res := make([]int, 0)
	columnMaxNumCache := make(map[int]int) // column index->relate column max num's row index


	for i, row := range matrix {
		rowMin := math.MaxInt32 // 每一行的最小值
		rowMinIndex := -1       // 每一行最小值的列 index
		columnMax := 0          // 每一行的最大值对应的列的最大值
		columnMaxRowIndex := 0  // 每一行的最大值对应的列的最大值的row index
		ok := false

		for j, num := range row {
			if num < rowMin {
				rowMin = num
				rowMinIndex = j
			}
		}

		columnMaxRowIndex, ok = columnMaxNumCache[rowMinIndex]
		if !ok {
			for i, row := range matrix {
				if row[rowMinIndex] > columnMax {
					columnMax = row[rowMinIndex]
					columnMaxRowIndex = i
				}
			}
			columnMaxNumCache[rowMinIndex] = columnMaxRowIndex
		}

		if i == columnMaxRowIndex {
			res = append(res, matrix[i][rowMinIndex])
		}

	}

	return res
}
