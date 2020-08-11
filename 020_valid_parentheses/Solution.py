class Solution:
    def isValid(self, s: str) -> bool:
        mapping = {')': '(', ']': '[', '}': '{'}
        stack = []
        for ch in s:
            if stack and stack[-1] == mapping.get(ch):
                stack.pop()
                continue
            stack.append(ch)

        return len(stack) == 0


if __name__ == '__main__':
    solution = Solution()
    print(solution.isValid('()[]{}'))

