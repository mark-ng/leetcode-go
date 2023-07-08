// Day 20

func combine(n int, k int) [][]int {
    global := [][]int{}
    temp := []int{}
    dfs(1, &temp, n, &global, k)
    return global 
}

func dfs(idx int, temp *[]int, n int, global *[][]int, k int) {
    if len(*temp) == k {
        tt := make([]int, 0, k) 
        for i := 0; i < k; i++ {
            tt = append(tt, (*temp)[i])
        } 
        *global = append(*global, tt)
        return
    }
    if idx == n + 1 {
        return
    } 
    
    *temp = append(*temp, idx)
    dfs(idx + 1, temp, n, global, k)
    *temp = (*temp)[:len(*temp) - 1]
    dfs(idx + 1, temp, n, global, k) 

}