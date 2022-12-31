from __future__ import print_function
from fractions import Fraction

def product(fracs):
    #input값 split해서 분자, 분모 map으로 받아서 유리수 만들어주는 Fraction 함수에 넣음
    #fracs는 [1/3, 3/4, 6/10]
    #reduce 람다에 곱해버림
    t = reduce(lambda x, y : x * y, fracs) # complete this line with a reduce statement
    return t.numerator, t.denominator

if __name__ == '__main__':
    fracs = []
    for _ in range(input()):
        fracs.append(Fraction(*map(int, raw_input().split())))
    result = product(fracs)
    print(*result)
