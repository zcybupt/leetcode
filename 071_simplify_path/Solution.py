class Solution:
    def simplifyPath(self, path: str) -> str:
        path_list = path.split('/')
        result = []
        for tmp in path_list:
            if tmp not in ['', '.', '..']:
                result.append(tmp)
            elif tmp == '..' and result:
                result.pop()

        return '/' + '/'.join(result)


if __name__ == '__main__':
    tests = [
        "/home/",
        "/../",
        "/home//foo/",
        "/a/../../b/../c//.//",
        "/a//b////c/d//././/.."
    ]

    solution = Solution()
    for test in tests:
        print(solution.simplifyPath(test))

