class Solution:
    def searchRange(self, nums: list, target: int) -> list:
        if not nums or len(nums) == 0: return [-1, -1]
        return [
            self.binary_search(nums, target, True),
            self.binary_search(nums, target, False)
        ]

    def binary_search(self, nums: list, target: int, is_search_first: bool):
        arr_len = len(nums)
        left, right = 0, arr_len - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] > target: right = mid - 1
            elif nums[mid] < target: left = mid + 1
            else:
                if is_search_first:
                    if mid > 0 and nums[mid] == nums[mid - 1]: right = mid - 1
                    else: return mid
                else:
                    if mid < arr_len - 1 and nums[mid] == nums[mid + 1]: left = mid + 1
                    else: return mid

        return -1


if __name__ == '__main__':
    nums = [5, 7, 7, 8, 8, 10]
    target = 8
    solution = Solution()
    print(solution.searchRange(nums, target))

