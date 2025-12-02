class Solution {

    private int a;
    private int b;

    private int [][] dp;

    private int solve(int i , int j , String s1 , String s2){
        if (i == a && j == b){
            return 0;
        }

        if (i >= a){
            return b - j; // add
        }

        if (j >= b){
            return a - i; // remove
        }
       
        if (dp[i][j] != -1){
            return dp[i][j];
        }
        int ret = Integer.MAX_VALUE;

        char char1 = s1.charAt(i);
        char char2 = s2.charAt(j);

        if (char1 == char2){
            ret = solve(i+1 , j +1 , s1 , s2);
        }else {

            // insert
            int ch1 = 1+ solve(i , j+1 , s1 , s2);


            // delete
            int ch2 = 1+solve(i+1 , j , s1 , s2);


            // replace 
            int ch3 = 1+solve(i+1 , j+1 , s1 , s2);

            ret = Math.min(Math.min(ret , ch1) , Math.min(ch2 , ch3));
        }
        return dp[i][j] = ret;
    }
    public int minDistance(String word1, String word2) {
        a = word1.length();
        b = word2.length();
        dp = new int[a+1][b+1];
        for (int i = 0; i <= a; i++){
            Arrays.fill(dp[i] , -1);
        }
        return solve(0 , 0 , word1 , word2);
        
    }
}