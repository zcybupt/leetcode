class Solution:
    def canJump(self, nums: list) -> bool:
        right_most = 0
        arr_len = len(nums)

        for i in range(arr_len):
            if i <= right_most:
                right_most = max(right_most, i + nums[i])
                if right_most > arr_len - 1: return True

        return False


if __name__ == '__main__':
    nums = [2, 3, 1, 1, 4]
    print(Solution().canJump(nums))

