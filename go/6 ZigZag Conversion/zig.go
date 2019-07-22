func convert(s string, numRows int) string {
    if numRows <= 1 {
        return s
    }
    ans := make([][]rune, numRows)
    direct := 1
    curr := 0
    for _, r := range s {
        if curr == 0 {
            direct = 1
        }
        if curr == numRows-1 {
            direct = -1
        }
        ans[curr] = append(ans[curr], r)
        curr += direct
    }
    ret := ""
    for _, i := range ans {
        for _, j := range i {
            ret += string(j)
        }
    }
    return ret
}
