class Solution {

    private int getValueOfChar(char c){
        return switch(c){
                case 'I' -> 1;
                case 'V' -> 5;
                case 'X' -> 10;
                case 'L' -> 50;
                case 'C' -> 100;
                case 'D' -> 500;
                default -> 1000;
            };
    }
    public int romanToInt(String s) {
        int ans = 0;
        char x = 'a';
        int n = s.length();
        for (int i = 0; i < n; i++){
            char c = s.charAt(i);
            ans+= getValueOfChar(c);
            if (i >0 && getValueOfChar(s.charAt(i-1)) < getValueOfChar(c)){
                ans-= 2 * getValueOfChar(s.charAt(i-1));
            }
        }
        return ans;
        
    }
}