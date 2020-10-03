class Solution:
    def minArray(self, nums: list) -> int:
        l, r = 0, len(nums) - 1

        while l < r:
            mid = (l + r) // 2
            if nums[mid] < nums[r]: r = mid
            elif nums[mid] > nums[r]: l = mid + 1
            else: r -= 1

        return nums[l]


if __name__ == '__main__':
    nums = [3, 4, 5, 1, 2]
    print(Solution().minArray(nums))
