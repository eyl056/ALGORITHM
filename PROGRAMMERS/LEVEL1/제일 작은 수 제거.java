import java.util.ArrayList;

class Solution {
    public int[] solution(int[] arr) {
        if (arr.length <= 1) {
            return new int[]{-1};
        }
        int min = arr[0];
        for (int i=0; i < arr.length; i++) {
            if (arr[i] < min) {
                min = arr[i];
            }
        }
        int answer[] = new int[arr.length - 1];
        int size = 0;
        for (int i: arr) {
            if (i != min) {
                answer[size++] = i;
            }
        }
        return answer;
    }
}
