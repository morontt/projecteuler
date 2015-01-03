# Integer right triangles
#
# Completed on Sun, 19 Oct 2014, 19:51

p_list = {}

x = 1
while x < 1000
    y = x + 1
    while y < 1001
        z = Math.sqrt (x * x + y * y)
        if z == Math.floor z
            p = x + y + z
            if p < 1001
                if p_list[p]?
                    p_list[p] += 1
                else
                    p_list[p] = 1
        y++
    x++

max = 0
max_p = 0
for own key, value of p_list
    if value > max
        max = value
        max_p = key

console.log max_p
