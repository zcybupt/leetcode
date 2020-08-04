import java.util.HashMap;

class FSM {
    public String status = "start";
    public int sign = 1;
    public long ans = 0;
    public HashMap<String, String[]> table = new HashMap<>();

    public FSM() {
        table.put("start", new String[]{"start", "signed", "in_number", "end"});
        table.put("signed", new String[]{"end", "end", "in_number", "end"});
        table.put("in_number", new String[]{"end", "end", "in_number", "end"});
        table.put("end", new String[]{"end", "end", "end", "end"});
    }

    public int getColumn(char ch) {
        if (ch == ' ') return 0;
        if (ch == '+' || ch == '-') return 1;
        if (ch >= '0' && ch <= '9') return 2;
        return 3;
    }

    public void get(char ch) {
        status = table.get(status)[getColumn(ch)];
        if (status.equals("in_number")) {
            int curInt = Character.getNumericValue(ch);
            ans = ans * 10 + curInt;
            if ((int) ans != ans) {
                if (sign == 1) ans = Integer.MAX_VALUE;
                if (sign == -1) ans = Integer.MIN_VALUE;
            }
        } else if (status.equals("signed")) {
            sign = ch == '+' ? 1 : -1;
        }
    }
}

class Solution {
    public int myAtoi(String str) {
        FSM fsm = new FSM();
        for (char ch : str.toCharArray()) {
            fsm.get(ch);
        }
        return fsm.sign * (int) fsm.ans;
    }
}
