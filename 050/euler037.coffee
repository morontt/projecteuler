# Truncatable primes
#
# Completed on Sun, 19 Oct 2014, 12:20

cache = []

is_prime = (t)->
    if t < 2
        return false
    if t is 2
        return true

    return cache[t] if cache[t]?

    m = Math.floor Math.sqrt t
    res = true
    i = 2

    while i <= m
        if t % i == 0
            res = false
            break
        i++

    cache[t] = res

    return res

truncatable_left = (t)->
    if t < 10
        return true

    tt = Math.floor (t / 10)

    return (is_prime tt) && (truncatable_left tt)

truncatable_right = (t)->
    if t < 10
        return true

    tt = parseInt t.toString().substring(1)

    return (is_prime tt) && (truncatable_right tt)

sum = 0
count = 0
n = 11
while count < 11
    if (is_prime n) and (truncatable_left n) and (truncatable_right n)
        sum += n
        count++
    n += 2

console.log sum
