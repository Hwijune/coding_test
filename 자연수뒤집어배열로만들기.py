#https://school.programmers.co.kr/learn/courses/30/lessons/12932

def solution(n):
    answer = []
    r_n = reversed(str(n))
    
    for i in r_n:
        answer.append(int(i))
    return answer
