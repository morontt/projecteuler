# Digit factorials
#
# Completed on Sun, 28 Sep 2014, 01:56

cache = []

factor = (n)->
    return cache[n] if cache[n]

    if n == 0
        result = 1
        cache[n] = result

        return result

    if n == 1 or n == 2
        result = n
    else
        result = n * factor (n - 1)

    cache[n] = result

    return result

digit_sum = (t)->
    s = t.toString()
    sum = 0
    i = 0
    len = s.length

    while i < len
        sum += factor parseInt s[i]
        i++

    return sum

j = 3
max = 6 * factor 9
r = 0

while j < max
    r += j if j is digit_sum j
    j++

console.log r
