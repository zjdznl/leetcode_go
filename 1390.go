package leetcode

import "math"

var allPrime = make([]bool, 100001)

// 素数筛选法
func init() {
	for i := range allPrime {
		allPrime[i] = true
	}
	allPrime[0] = false
	allPrime[1] = false

	for i := 2; i < len(allPrime); i++ {
		if allPrime[i] == true {
			j := i
			for {
				product := i * j
				if product >= len(allPrime) {
					break
				}
				allPrime[product] = false
				j++
			}
		}
	}
}

func sumFourDivisors(nums []int) int {
	sum := 0

	for _, n := range nums {
		sqrt := int(math.Sqrt(float64(n)))
		for i := 2; i <= sqrt; i++ {
			if allPrime[i] && n%i == 0 { // i is a prime
				divisor := n / i
				if allPrime[divisor] || divisor == i*i{ // n is a product of two prime or is cube of a prime
					if i != divisor {
						sum = sum + 1 + n + i + divisor
					}
				}
			}
		}
	}

	return sum
}
