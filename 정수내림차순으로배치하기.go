import "strconv"
import "sort"
import "math"

func solution(n int64) int64 {
    var n_str string
    var n_slice[] int64
	var res int64
	
    n_str = strconv.FormatInt(n, 10)
    
	for i:=0 ; i<len(n_str); i++{
        n,_ = strconv.ParseInt(string(n_str[i]), 10, 64)
		n_slice = append(n_slice, n)
    }
	sort.Slice(n_slice, func(i, j int) bool { return n_slice[i] < n_slice[j] })
	
	for i:=0; i<len(n_slice); i++{
		res = res + int64(n_slice[i])*int64(math.Pow(float64(10), float64(i)))
	}
	
	return res
}
