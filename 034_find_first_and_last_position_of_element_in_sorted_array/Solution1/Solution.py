class Solution:
    def searchRange(self, nums: list, target: int) -> list:
        if not nums or len(nums) == 0: return [-1, -1]
        arr_len = len(nums)
        left , right = 0, arr_len - 1

        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target:
                while nums[left] != target: left += 1
                while nums[right] != target: right -= 1
                return [left, right]
            if nums[mid] < target: left = mid + 1
            else: right = mid - 1

        return [-1, -1]


if __name__ == '__main__':
    nums = [5, 7, 7, 8, 8, 10]
    target = 8
    solution = Solution()
    print(solution.searchRange(nums, target))

