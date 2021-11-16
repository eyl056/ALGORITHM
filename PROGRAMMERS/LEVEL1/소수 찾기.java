import java.util.Arrays;
import java.util.ArrayList;
                        
class Solution {
    public int solution(int n) {
        boolean sieve[] = new boolean[n + 1];
        Arrays.fill(sieve, true);
        int m = (int) Math.sqrt(n);
        for (int i = 2; i < m + 1; i++) {
            if (sieve[i] == true) {
                for (int j = i + i; j < n + 1; j += i) {
                    sieve[j] = false;
                }
            }
        }
        
        ArrayList<Integer> answer = new ArrayList<Integer>();
        for (int i = 2;i < n + 1; i++) {
            if (sieve[i] == true) {
                answer.add(i);
            }
        }
        return answer.size();
    }
}
