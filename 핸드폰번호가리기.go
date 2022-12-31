#https://programmers.co.kr/learn/courses/30/lessons/12948

func solution(phone_number string) string {
    var str_len int
    str_len = len(phone_number)

    var res string
    
    for i:= 0 ; i < str_len ; i++{
        if i < str_len-4 {
            res = res + "*"
        } else {
            res = res + string(phone_number[i])
        }
    }
    return res
}
