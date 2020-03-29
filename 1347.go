package leetcode

func minSteps(s string, t string) int {
	count := make([]int, 26)
	length := len(s)
	for i := 0; i < length; i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	res := 0
	for _, c := range count {
		if c > 0 {
			res += c
		}
	}
	return res
}
