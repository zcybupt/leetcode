class Solution:
    def search(self, nums: list, target: int) -> int:
        arr_len = len(nums)
        left, right = 0, arr_len - 1

        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target: return mid
            if nums[mid] < nums[right]: # 右侧有序
                if nums[mid] < target <= nums[right]: left = mid + 1
                else: right = mid - 1
            else:
                if nums[left] <= target < nums[mid]: right = mid - 1
                else: left = mid + 1

        return -1


if __name__ == '__main__':
    nums = [4, 5, 6, 7, 0, 1, 2]
    target = 0
    solution = Solution()
    print(solution.search(nums, target))
