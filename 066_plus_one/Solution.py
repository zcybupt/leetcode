class Solution:
    def plusOne(self, digits: list) -> list:
        arr_len = len(digits)
        for i in range(arr_len - 1, -1, -1):
            if digits[i] != 9:
                digits[i] += 1
                return digits
            digits[i] = 0

        return [1] + [0] * arr_len


if __name__ == '__main__':
    digits = [9, 9, 9, 9]
    print(Solution().plusOne(digits))
