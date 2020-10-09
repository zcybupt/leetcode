class Solution:
    def lengthOfLIS(self, nums: list) -> int:
        arr_len = len(nums)
        if arr_len < 2: return arr_len

        tails = [0] * arr_len
        result = 0

        for num in nums:
            l, r = 0, result
            while l < r:
                m = (l + r) // 2
                if tails[m] < num: l = m + 1
                else: r = m
            tails[l] = num
            if result == r: result += 1

        return result


if __name__ == '__main__':
    nums = [4, 10, 4, 3, 8, 9]
    print(Solution().lengthOfLIS(nums))

