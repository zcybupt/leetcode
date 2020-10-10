import java.util.*;

class MinStack {
    List<Integer> stackA, stackB;
    int lastA = -1, lastB = -1;

    public MinStack() {
        stackA = new ArrayList<>();
        stackB = new ArrayList<>();
    }

    public void push(int x) {
        stackA.add(x);
        lastA++;
        if (stackB.isEmpty() || stackB.get(lastB) >= x) {
            stackB.add(x);
            lastB++;
        }
    }

    public void pop() {
        if (stackA.get(lastA).equals(stackB.get(lastB))) {
            stackB.remove(lastB);
            lastB--;
        }
        stackA.remove(lastA);
        lastA--;
    }

    public int top() {
        return stackA.get(lastA);
    }

    public int getMin() {
        return stackB.get(lastB);
    }
}

class Solution {
    public static void main(String[] args) {
        MinStack stack = new MinStack();
        stack.push(-2);
        stack.push(0);
        stack.push(-3);
        System.out.println(stack.getMin());
        stack.pop();
        System.out.println(stack.top());
        System.out.println(stack.getMin());
    }
}
