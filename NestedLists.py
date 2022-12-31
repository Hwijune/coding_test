if __name__ == '__main__':
    records = [] #[이름, 점수] 이중리스트
    set_list = [] #중복제거 리스트
    for _ in range(int(input())):
        name = input()
        score = float(input())
        records.append([name, score])
        set_list.append(score)
    
    set_list = list(set(set_list))
    set_list.sort() 리스트 두번째 값
    
    records.sort(key=lambda x: (x[1], x[0])) #이름 알파벳순으로 정렬
    second_lowest = set_list[1]
    for num in records:
        if(num[1] == second_lowest):
            print(num[0]) 
