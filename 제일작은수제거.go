//https://programmers.co.kr/learn/courses/30/lessons/12935

func solution(arr []int) []int {
	arr_len := len(arr)
	var res []int
	
	minvalue := 1000000
	 
	for i, e := range arr {
		if i==0 || e < minvalue {
			minvalue = e
		}
	}
    for i:=0; i<arr_len; i++ {
		if minvalue != arr[i] {
			res = append(res, arr[i])
		}
	}

    if arr_len == 1 {
        return []int{-1}    
    }
	return res
}
