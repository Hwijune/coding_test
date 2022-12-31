#https://programmers.co.kr/learn/courses/30/lessons/12954

func solution(x int, n int) []int64 {
    var s []int64
    
    for i:= 0 ; i<n; i++ {
        s = append(s, int64(x)*(int64(i)+1))
    }
    return s
}
