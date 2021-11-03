import java.util.Arrays;

class Solution {
    public int[] solution(long n) {
        String ns = Long.toString(n);
        int[] answer = new int[ns.length()];
        
        int i = 0;
        while (n > 0) {
            answer[i] = (int)(n % 10);
            n /= 10;
            i++;
        }
        return answer;
    }
}
