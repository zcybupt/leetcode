class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        # if s.strip() == '': return 0
        # return len(s.split()[-1])
        length = 0

        for i in range(len(s) - 1, -1, -1):
            if s[i] == ' ':
                if length > 0: break
                else: continue
            length += 1

        return length


if __name__ == '__main__':
    print(Solution().lengthOfLastWord('Hello World '))

