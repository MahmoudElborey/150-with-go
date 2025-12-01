class Solution {
    private int a;
    private int b;
    private int [][][] dp;
    // 1 can be done
    // -1 didn't visit this state yet
    // 0 can't be done
    private int solve(int flag , int i , int j ,String s1 , String s2 , String s3){
        System.out.println("flag is " + flag + " i is " + i + " j is " + j);
        if (i == a && b == j){
            return 1;
        }


        if (dp[flag][i][j] != -1){
            return dp[flag][i][j];
        }

        dp[flag][i][j] = 0;
        int ch = 0;
       
            for (int k = (flag == 0 ? i : j); k < (flag == 0 ? a : b); ++k){
                int curIdx = k + (flag == 0 ? j : i);
                if (flag == 0 && s1.charAt(k) == s3.charAt(curIdx)){
                    ch = Math.max(ch , solve(flag^1 , k+1 , j , s1 , s2 , s3));
                    
                }else if (flag == 1 && s2.charAt(k) == s3.charAt(curIdx)){
                     ch = Math.max(ch , solve(flag^1 , i , k+1 , s1 , s2 , s3));
                }else {
                    break;
                }
            }


        return dp[flag][i][j] = ch;
    }
    public boolean isInterleave(String s1, String s2, String s3) {
        a = s1.length();
        b = s2.length();
        int c = s3.length();
        if (c != a+b){
            return false;
        }
        dp = new int[2][a+1][b+1];
        for (int i = 0; i < 2; ++i){
            for (int j = 0; j <= a; ++j){
                for (int k = 0; k <= b; ++k){
                    dp[i][j][k] = -1;
                }
            }
        }
        int ans = 0;
        for (int i = 0; i < 2; ++i){
            ans = Math.max(ans , solve(i , 0 , 0 , s1 , s2 , s3));
        }
        return ans == 1 ? true : false;

    }
}