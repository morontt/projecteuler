# Convergents of e
#
# Completed on Sun, 29 Nov 2015, 22:55

n = 100
k = (n - 1) // 3 + 2

dlist = [2]

for i in range(1, k):
    dlist.append(1)
    dlist.append(2 * i)
    dlist.append(1)

a = 1
b = dlist[n - 1]

j = n - 2
while j >= 0:
    a, b = b, b * dlist[j] + a
    j -= 1

summ = 0
for s in list(str(b)):
    summ += int(s)

print summ
