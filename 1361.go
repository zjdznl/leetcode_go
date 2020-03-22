package leetcode

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	sideNum := 0
	for _, v := range append(leftChild, rightChild...) {
		if v != -1 {
			sideNum += 1
		}
	}
	return sideNum == n-1
}
