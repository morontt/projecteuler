# Champernowne's constant
#
# Completed on Sun, 12 Oct 2014, 12:19

n = 1
s = 1
idx = [1, 10, 100, 1000, 10000, 100000, 1000000]

result = 1

while n < 7
    min = Math.pow 10, (n - 1)
    max = (Math.pow 10, n) - 1
    s_old = s
    s += (max - min + 1) * n

    for i in idx
        if i >= s_old and i < s
            delta = Math.floor ((i - s_old) / n)
            str = (min + delta).toString()
            result *= parseInt(str[(i - s_old) % n])
    n++

console.log(result)
