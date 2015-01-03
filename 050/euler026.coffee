# Reciprocal cycles
#
# Completed on Sat, 11 Oct 2014, 23:58

divide = (t)->
    rest = 1
    cycle = {}
    res = 0
    i = 0

    while rest != 0
        i++
        if rest >= t
            rest %= t
        rest *= 10

        if cycle[rest]?
            res = i - cycle[rest]
            rest = 0
        else
            cycle[rest] = i

    return res

max = 0
max_d = 0

d = 2
while d < 1000
    t = divide d
    if max < t
        max = t
        max_d = d
    d++

console.log max_d
