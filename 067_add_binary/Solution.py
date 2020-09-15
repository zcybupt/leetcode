class Solution:
    def addBinary(self, a: str, b: str) -> str:
        x, y = int(a, 2), int(b, 2)

        def add(x, y):
            if not y: return x
            xor = x ^ y
            carry = (x & y) << 1
            return add(xor, carry)

        res = add(x, y)
        return bin(res)[2:]


if __name__ == '__main__':
    print(Solution().addBinary('11', '1'))

