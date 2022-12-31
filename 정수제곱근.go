#https://school.programmers.co.kr/learn/courses/30/lessons/12934?language=go

import "math"

func solution(n int64) int64 {
    var m_sqrt int64 
    m_sqrt = int64(math.Sqrt(float64(n)))
    var res int64
    res = -1
    
    if m_sqrt * m_sqrt == n {
        res = (m_sqrt+1)*(m_sqrt+1)
    }
    
    return res
}
