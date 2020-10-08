class Solution:
    def lengthOfLIS(self, nums: list) -> int:
        arr_len = len(nums)
        if arr_len < 2: return arr_len

        dp = [1] * arr_len

        for i in range(1, arr_len):
            for j in range(i):
                if nums[i] > nums[j]:
                    dp[i] = max(dp[i], dp[j] + 1)

        return max(dp)


if __name__ == '__main__':
    nums = [10, 9, 2, 5, 3, 7, 101, 18]
    print(Solution().lengthOfLIS(nums))
