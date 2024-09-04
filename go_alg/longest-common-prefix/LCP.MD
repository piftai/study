Algorithm:

1. The longest prefix will be from the first word, it means that your start point is first word.
2. Next step is run every step of slice and compare if first word or his part(?) is equal to word
of slice or his part that have to be the same length
3. More information. Step By Step algorithm
4. prefix := strs[0]
for i := len(prefix)-1; i != 0; i-- {
    var cur string
    j, v := range strs {
        cur = v
        if v[i:] == prefix[i:] && len(v) >= i{
            cnt++
        } else {
            continue
        }
        if cnt == len(strs) { break; return curr }
   }
}