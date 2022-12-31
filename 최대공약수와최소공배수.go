#https://programmers.co.kr/learn/courses/30/lessons/12940

func euclid(a int, b int) int{
    if b == 0 {
        return a
    } else {
        return euclid(b, a%b)
    }
}
func solution(n int, m int) []int {
    euc := euclid(n, m)
    return []int{euc, n*m / euc}
}
