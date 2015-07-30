# Passcode derivation
#
# Completed on 29 Jul 2015, 02:14

auths = []
symbols = ['0','1','2','3','6','7','8','9']

f = open('./p079_keylog.txt', 'r')
for line in f:
    auths.append(line.strip())

f.close()

def to_str(x):
    s = ''
    o = str(oct(x))
    for i in o[1:]:
        s += symbols[int(i)]

    return s

def auth_check(key, code):
    res = False
    i1 = key.find(code[0])
    if i1 > -1:
        i2 = key.find(code[1], i1)
        if i2 > -1:
            i3 = key.find(code[2], i2)
            if i3 > -1:
                res = True

    return res

test = 2097151
ll = len(auths)
t = 0

while t < ll:
    s = to_str(test)
    t = 0
    for code in auths:
        if auth_check(s, code):
            t += 1
        else:
            break
    test += 1

print s
