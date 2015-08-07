# Cubic permutations
#
# Completed on 5 Aug 2015, 01:39

def p_hash(x):
    arr = list(str(x))
    arr.sort()
    s = ''

    return s.join(arr)

target = 5
n = 3
find = True
result = []

while find:
    rmin = int((10 ** n) ** (1 / 3.0)) + 1
    rmax = int((10 ** (n + 1) - 1) ** (1 / 3.0))
    d = {}
    for i in xrange(rmin, rmax + 1):
        cube = i ** 3
        cube_hash = p_hash(cube)
        if cube_hash in d:
            t = d[cube_hash]
            t.append(cube)
        else:
            d[cube_hash] = [cube]

    for key in d:
        if len(d[key]) == target:
            find = False
            result += d[key]

    n += 1

print min(result)
