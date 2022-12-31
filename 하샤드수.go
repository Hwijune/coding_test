#https://programmers.co.kr/learn/courses/30/lessons/12947

import "strconv"

func solution(x int) bool {
    var res bool
    var x_str string
    var x_sum int
    
    x_str = strconv.Itoa(x)
	//println(x_str)
	
    for i:=0 ; i<len(x_str); i++{
		b, _ := strconv.Atoi(string(x_str[i]))
		x_sum = x_sum + b
    }
    
	//println(x_sum)
	
    if x%x_sum == 0 {
        res = true
    }
    return res
}
