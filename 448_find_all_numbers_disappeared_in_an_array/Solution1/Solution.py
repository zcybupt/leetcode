class Solution:
    def findDisappearedNumbers(self, nums: list) -> list:
        if not nums: return []

        result = []
        tmp_arr = [0] * len(nums)
        for num in nums: tmp_arr[num - 1] += 1
        for i, num in enumerate(tmp_arr):
            if num == 0: result.append(i + 1)

        return result

    def findDisappearedNumbers2(self, nums: list) -> list:
        if not nums: return []

        tmp_dict = {}
        for num in nums: tmp_dict[num] = tmp_dict.get(num, 0) + 1
        result = []
        for i in range(1, len(nums) + 1):
            if i not in tmp_dict: result.append(i)

        return result



if __name__ == '__main__':
    nums = [4, 3, 2, 7, 8, 2, 3, 1]
    print(Solution().findDisappearedNumbers2(nums))

