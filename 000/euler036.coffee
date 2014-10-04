# Double-base palindromes
#
# Completed on Sat, 4 Oct 2014, 18:37

to_bin = (x)->
  res = ''
  while x > 0
    res += x % 2
    x = x >> 1

  return res

is_palindrome = (t)->
  len = t.length

  if len is 0 or len is 1
    return true

  t[0] is t[len - 1] and is_palindrome t.substr(1, len - 2)

limit = 1000000
i = 1
sum = 0

while i < limit
  if is_palindrome i.toString()
    k = to_bin i
    if is_palindrome k
      sum += i

  i++

console.log sum
