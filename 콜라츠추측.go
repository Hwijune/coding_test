#https://programmers.co.kr/learn/courses/30/lessons/12943#
func solution(num int) int {
    var cnt int
    cnt = 0
	
    for cnt != 500 && num != 1{
		if num%2 == 0 {
            num = num/2
		} else {
            num = num*3 + 1
            
        }
		cnt = cnt + 1
    }

    if cnt == 500 {
        return -1
    }
    return cnt
}
