class Solution {
    public String solution(String s, int n) {
        String answer = "";
        for (int i = 0; i < s.length(); i++) {
            char ss = s.charAt(i);
            if (Character.isUpperCase(ss)) {
                ss = (char)((ss - 'A' + n) % 26 + 'A');
            }
            else if (Character.isLowerCase(ss)) {
                ss = (char)((ss - 'a' + n) % 26 + 'a');
            }
            answer += ss;
        }
        return answer;
    }
}
