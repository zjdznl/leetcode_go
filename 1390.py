import math

all_prime = [True] * 100001
all_prime[0] = False
all_prime[1] = False
has_called = False


class Solution(object):
    def __init__(self):
        global has_called
        if has_called:
            return

        for i in range(2, len(all_prime)):
            if all_prime[i]:
                j = i
                while True:
                    product = i * j
                    if product >= len(all_prime):
                        break
                    all_prime[product] = False
                    j += 1
        has_called = True

    def sumFourDivisors(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        sum = 0
        for n in nums:
            sqrt = int(math.sqrt(n))
            for i in range(2, sqrt + 1):
                if all_prime[i] and n % i == 0:
                    divisor = int(n / i)
                    if all_prime[divisor] or divisor == i * i:
                        if i != divisor:
                            sum = sum + 1 + n + i + divisor
                            break

        return sum


if __name__ == '__main__':
    s = Solution()
    print(s.sumFourDivisors(list(range(1, 11))))
