class FSM:
    def __init__(self):
        self.state = 'start'
        self.sign = 1
        self.ans = 0
        self.table = {
            'start': ['start', 'signed', 'in_number', 'end'],
            'signed': ['end', 'end', 'in_number', 'end'],
            'in_number': ['end', 'end', 'in_number', 'end'],
            'end': ['end', 'end', 'end', 'end'],
        }

    def get_column(self, c: str) -> int:
        if c.isspace():
            return 0
        if c in ['+', '-']:
            return 1
        if c.isdigit():
            return 2
        return 3

    def get(self, c: str):
        self.state = self.table[self.state][self.get_column(c)]
        if self.state == 'in_number':
            self.ans = self.ans * 10 + int(c)
            self.ans = min(self.ans, INT_MAX) if self.sign == 1 else min(self.ans, -INT_MIN)
        elif self.state == 'signed':
            self.sign = 1 if c == '+' else -1


class Solution:
    def myAtoi(self, str: str) -> int:
        fsm = FSM()
        for c in str:
            fsm.get(c)
        return fsm.sign * fsm.ans
