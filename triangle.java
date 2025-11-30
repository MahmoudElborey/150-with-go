class Solution {
    public int minimumTotal(List<List<Integer>> triangle) {
        int rows = triangle.size();
        int [] dp = new int[rows];
        for (int row = rows-1; row >= 0; --row){
            int curLength = triangle.get(row).size();
            int [] newDp = new int[curLength];

            for (int col = curLength -1; col >= 0; --col){
                int value = triangle.get(row).get(col);
                if (row+1 == rows){
                    newDp[col] = value;
                }else {
                    newDp[col] = value + Math.min(dp[col] , dp[col+1]);
                }
                System.out.println("cur Row " + row + " cur col " + col + " value  " + newDp[col]);
            }

            dp = newDp;

        }

        return dp[0];

        
        
    }
}