import java.lang.Math;

class Solution {
    public long solution(long n) {
        if ((int)Math.sqrt(n) == Math.sqrt(n)) {
            return (long) Math.pow((Math.sqrt(n) + 1), 2);
        }
        
        return -1;
    }
}
