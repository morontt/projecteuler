# XOR decryption
#
# Completed on Thu, 30 Jul 2015, 01:56

f = open('./p059_cipher.txt', 'r')
tmp = f.readline().strip().split(',')
f.close()

data = map(lambda c: int(c), tmp)


def cipher_decode(arr, key):
    decoded = []
    i = 0

    for c in arr:
        decoded.append(chr(c ^ ord(key[i % 3])))
        i += 1

    return decoded

"""
run:
    python 075/euler059.py | grep -i word

password: god
"""

for k1 in xrange(97, 123):
    for k2 in xrange(97, 123):
        for k3 in xrange(97, 123):
            passwd = chr(k1) + chr(k2) + chr(k3)
            z = cipher_decode(data, passwd)
            s = ''
            print passwd + '---' + s.join(z)

sum0 = 0
dec = cipher_decode(data, 'god')

for c in dec:
    sum0 += ord(c)

print sum0
