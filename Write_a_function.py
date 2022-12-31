'''
https://www.hackerrank.com/challenges/write-a-function/problem?isFullScreen=true&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

윤년구하기인데 문제가 설명이 엄청 길어져서 처음에 당황했다
'''

def is_leap(year):
    leap = False
    # Write your logic here
    if (year % 4 == 0):
        leap = True
        if (year % 100 == 0):
            leap = False
        if (year % 400 == 0):
            leap = True
    return leap

year = int(raw_input())
print is_leap(year)
