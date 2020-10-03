class Solution:
    def search(self, nums: list, target: int) -> int:
        if not nums: return -1

        l, r = 0, len(nums) - 1
        while l <= r:
            mid = (l + r) // 2
            if nums[mid] == target: return mid
            elif nums[mid] > target: r = mid - 1
            else: l = mid + 1

        return -1


if __name__ == '__main__':
    nums = [-1, 0, 3, 5, 9, 12]
    print(Solution().search(nums, 9))

