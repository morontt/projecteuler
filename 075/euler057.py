# Square root convergents
#
# Completed on Fri, 31 Jul 2015, 14:20

a0, b0 = 3, 2
i = 1
r = 0

while i < 1000:
    a1, b1 = a0 + 2 * b0, a0 + b0
    if len(str(a1)) > len(str(b1)):
        r += 1
    a0, b0 = a1, b1
    i += 1

print r
