# Quadratic primes
#
# Completed on Mon, 6 Oct 2014, 23:33

cache = []
limit = 1000

max_length = 0;
max_a = 0
max_b = 0

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

b_cycle = (w)->
    a = -w
    while a < limit
        n = 0
        z = n * n + a * n + w
        while is_prime z
            n++
            z = n * n + a * n + w

        if n > max_length
            max_length = n
            max_a = a
            max_b = w

        a += 2

    return false

b_cycle 2

b = 3
while b < limit
    b_cycle b
    b += 2

console.log max_a * max_b
