class Solution:
    def maxSubArray(self, nums):
        if not nums or len(nums) == 0: return 0

        current = 0
        max_sum = float('-inf')

        for num in nums:
            current += num
            max_sum = max(current, max_sum)
            current = max(current, 0)

        return max_sum


if __name__ == '__main__':
    print(Solution().maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]))

