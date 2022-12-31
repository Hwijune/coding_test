#https://ide-run.goorm.io/workspace/dG1Ly5KaO4eYTe9hynw?language=kor

func solution(arr []int) float64 {
    var sum float64
    
    for i:=0; i<len(arr); i++{
        sum += float64(arr[i])
    }
    return sum/float64(len(arr))
}
