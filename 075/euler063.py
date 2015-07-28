# 28.07.2015 14:40:45

def nlength(x, n):
    return len(str(x ** n))

res = 0
i = 1
while nlength(9, i) == i:
    res += 1
    j = 8
    while j > 0 and nlength(j, i) == i:
        res += 1
        j -= 1
    i += 1

print res
