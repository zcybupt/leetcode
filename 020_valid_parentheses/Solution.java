import java.util.Stack;

public class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        for (char ch : s.toCharArray()) {
            if (!stack.isEmpty() && stack.peek() == getPair(ch)) {
                stack.pop();
                continue;
            }
            stack.add(ch);
        }

        return stack.isEmpty();
    }

    public char getPair(char right) {
        switch (right) {
            case ')': return '(';
            case ']': return '[';
            case '}': return '{';
            default: return '#';
        }
    }

    public boolean isValid2(String s) {
        int sLen = s.length();
        if (sLen % 2 != 0) return false;
        char[] stack = new char[sLen];
        int index = 0;
        for (char ch : s.toCharArray()) {
            if (index > 0 && stack[index-1] == getPair(ch)) {
                index--;
                continue;
            }
            stack[index++] = ch;
        }

        return index == 0;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.isValid2("()[{}]"));
    }
}
