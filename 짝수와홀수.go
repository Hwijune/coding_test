#https://programmers.co.kr/learn/courses/30/lessons/12937

func solution(num int) string {
    var res string
    
    if num%2 == 0{
        res = "Even"
    } else {
        res = "Odd"
    }
    return res
}
