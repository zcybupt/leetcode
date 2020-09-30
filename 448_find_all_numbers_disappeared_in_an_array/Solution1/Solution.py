class Solution:
    def findDisappearedNumbers(self, nums: list) -> list:
        if not nums: return []
        tmp_arr = [0] * len(nums)
        for num in nums: tmp_arr[num - 1] += 1
        return [i + 1 for i, num in enumerate(tmp_arr) if num == 0]

if __name__ == '__main__':
    nums = [4, 3, 2, 7, 8, 2, 3, 1]
    print(Solution().findDisappearedNumbers(nums))

