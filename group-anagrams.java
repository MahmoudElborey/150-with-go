class Solution {

    private String sortString(String s){
        char [] chars = s.toCharArray();
        Arrays.sort(chars);
        return new String(chars);
    }
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String , Integer > mp = new HashMap<>();

        List<List<String>> ans = new ArrayList<>();

        for (String s : strs){
            String sorted = sortString(s);
            
            int idx = mp.getOrDefault(sorted , -1);

            if (idx == -1){
                idx = ans.size();
                mp.put(sorted , idx);
                ans.add(new ArrayList<>());
            
            }

           ans.get(idx).add(s);
        }

        return ans;
        
    }
}