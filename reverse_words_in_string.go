package withgo

import "strings"

func reverseWords(s string) string {
    s = strings.TrimSpace(s)
    words:= strings.Fields(s)
    n:= len(words)
    var sb strings.Builder
    for i:= n-1; i >= 0; i-- {
        if i !=  n-1{
          sb.WriteString(" ")
        }
        sb.WriteString(words[i])
    }

    return sb.String()
}