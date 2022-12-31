if __name__ == '__main__':
    x = int(input())
    y = int(input())
    z = int(input())
    n = int(input())
    
    res = "["
    for xi in range(x+1):
        for yi in range(y+1):
            for zi in range(z+1):
                if(xi+yi+zi != n):
                    res = res + "["+str(xi)+", "+str(yi)+", "+str(zi)+"], "
    if(len(res)>2):
        res = res[:-2]+"]"
    else:
        res = res+"]"
    print(res)
            
