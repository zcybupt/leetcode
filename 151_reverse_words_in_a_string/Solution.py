class Solution:
    def reverseWords(self, s: str) -> str:
        s = s.strip()
        l = r = len(s) - 1
        result = []
        while l >= 0:
            while l >= 0 and s[l] != ' ': l -= 1
            result.append(s[l + 1: r + 1])
            while l >= 0 and s[l] == ' ': l -= 1
            r = l
        return ' '.join(result)


if __name__ == '__main__':
    string = 'Stay hungry, stay foolish.'
    print(Solution().reverseWords(string))

