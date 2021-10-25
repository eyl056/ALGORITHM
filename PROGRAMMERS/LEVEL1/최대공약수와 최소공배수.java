class Solution {
    public int[] solution(int n, int m) {
        int[] answer = new int [2];
        int x = n;
        int y = m;
        while (m > 0) {
            int preN = n;
            n = m;
            m = preN % m;
        }
        answer[0] = n;
        answer[1] = x * y / answer[0];
        
        return answer;
    }
}
