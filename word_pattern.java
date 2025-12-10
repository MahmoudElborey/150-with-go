class Solution {
    public boolean wordPattern(String pattern, String s) {
        Map<String, Character> mp = new HashMap<>();
        String [] strings = s.split(" ");
        Map<Character , String> chars = new HashMap<>();
        int j = 0;

        if (strings.length != pattern.toCharArray().length) return false; 

        for (char cc : pattern.toCharArray()){
            boolean ch1 = mp.containsKey(strings[j]);
            boolean ch2 = chars.containsKey(cc);

            if (ch1 || ch2){
                if (!(ch1 && ch2 && mp.get(strings[j]).equals(cc) && chars.get(cc).equals(strings[j]))){
                    return false;
                }
            
            }else {
                mp.put(strings[j] , cc);
                chars.put(cc , strings[j]);
            }
            j++;
        }
        return true;
        
    }
}