/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {

    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> ans = new ArrayList<>();
        if (root == null){
            return ans;
        }

       Deque<TreeNode> queue = new ArrayDeque<>();

       queue.addLast(root);
        while(!queue.isEmpty()){
           int size = queue.size();
            while(size-- > 0){
                TreeNode top = queue.poll();
                if (size == 0){
                   ans.add(top.val);
                }
                if (top.left != null){
                    queue.addLast(top.left);
                }
                if (top.right != null){
                    queue.addLast(top.right);
                }
               
          }
        }
        return ans;

        
    }
}