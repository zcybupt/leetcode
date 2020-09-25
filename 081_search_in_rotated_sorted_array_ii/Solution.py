class Solution:
    def search(self, nums: list, target: int) -> bool:
        # return target in nums
        arr_len = len(nums)
        l, r = 0, arr_len - 1

        while l <= r:
            while l < r and nums[l] == nums[l + 1]: l += 1
            while l < r and nums[r] == nums[r - 1]: r -= 1
            mid = (l + r) // 2
            if nums[mid] == target: return True
            if nums[mid] < nums[r]:
                if nums[mid] < target <= nums[r]: l = mid + 1
                else: r = mid - 1
            else:
                if nums[l] <= target < nums[mid]: r = mid - 1
                else: l = mid + 1

        return False


if __name__ == '__main__':
    nums = [4, 5, 6, 7, 0, 1, 2]
    print(Solution().search(nums, 2))

