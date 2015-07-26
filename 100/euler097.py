# Large non-Mersenne prime
#
# Completed on Mon, 27 Jul 2015, 00:07

def to_binary(x):
    res = []
    i = 0
    while x > 0:
        if x % 2 == 1:
            res.append(i)
        i += 1
        x = (x - x % 2) / 2

    return res

def split_production(x, y):

    return (x * y) % 10000000000

def square_loop(t):
    res = 2
    while t > 0:
        res = split_production(res, res)
        t -= 1

    return res

tmp = 28433
for i in to_binary(7830457):
    tmp = split_production(tmp, square_loop(i))

print tmp + 1
