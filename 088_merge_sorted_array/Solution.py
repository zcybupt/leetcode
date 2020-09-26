class Solution:
    def merge(self, nums1: list, m: int, nums2: list, n: int) -> None:
        idx = m + n - 1
        m, n = m - 1, n - 1
        while m >= 0 and n >= 0:
            if nums1[m] > nums2[n]:
                nums1[idx] = nums1[m]
                idx, m = idx - 1, m - 1
            else:
                nums1[idx] = nums2[n]
                idx, n = idx - 1, n - 1

        while n >= 0:
            nums1[idx] = nums2[n]
            idx, n = idx - 1, n - 1


if __name__ == '__main__':
    nums1 = [1, 2, 3, 0, 0, 0]
    nums2 = [2, 5, 6]
    m, n = 3, 3
    Solution().merge(nums1, m, nums2, n)
    print(nums1)

