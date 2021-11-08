// 공백 공백 공백!!! *split(" ", -1), trim(X)*
import java.util.*;

class Solution {
    public String solution(String s) {
        String answer[] = s.split(" ", -1);
        String strAnswer = "";
        
        for (int i = 0;i < answer.length; i++) {
            String str = answer[i];
            for (int j = 0; j < str.length(); j++) {
                if (j % 2 == 0) {
                    strAnswer += Character.toUpperCase(str.charAt(j));
                }
                else {
                    strAnswer += Character.toLowerCase(str.charAt(j));
                }
            }
            if (i == answer.length - 1) {break;}
            strAnswer += " ";
        }
        return strAnswer;
    }
}
