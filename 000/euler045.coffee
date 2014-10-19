# Triangular, pentagonal, and hexagonal
#
# Completed on Sun, 19 Oct 2014, 15:22

n = 143
res = true

while res
    n++
    p = (1 + Math.sqrt(48 * n * n - 24 * n + 1)) / 6
    if p == Math.floor p
        res = false

console.log p * (3 * p - 1) / 2
