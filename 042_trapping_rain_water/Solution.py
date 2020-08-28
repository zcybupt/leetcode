class Solution:
    def trap(self, height_arr):
        if not height_arr or len(height_arr) < 3: return 0

        arr_len = len(height_arr)
        left_max, right_max = height_arr[0], height_arr[arr_len - 1]
        l, r = 1, arr_len - 2
        result = 0

        while l <= r:
            if left_max <= right_max:
                result += max(0, left_max - height_arr[l])
                left_max = max(left_max, height_arr[l])
                l += 1
            else:
                result += max(0, right_max - height_arr[r])
                right_max = max(right_max, height_arr[r])
                r -= 1

        return result


if __name__ == '__main__':
    height_arr = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
    solution = Solution()
    print(solution.trap(height_arr))

