import java.util.Arrays;

class Solution {
    public long solution(long n) {
        long[] arr = new long [(int)(Math.log10(n) + 1)];
        int i = 0;
        while(n > 0) {
            arr[i] = n % 10;
            n /= 10;
            i++;
        }
        Arrays.sort(arr);
        long answer = 0;
        int x = 1;
        for (long y : arr) {
            answer += y * x;
            x *= 10;
        }
        return answer;
    }
}
