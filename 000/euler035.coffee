# Circular primes
#
# Completed on Sat, 4 Oct 2014, 22:59

list = [2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, 97]
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

array_unique = (arr)->
  res = []
  test = []
  for u in arr
    unless test[u]?
      test[u] = true
      res.push u

  return res

i0 = 0
while i0 < 10
  i1 = 0
  while i1 < 10
    i2 = 0
    while i2 < 10
      x0 = i0 * 100 + i1 * 10 + i2
      x1 = i1 * 100 + i2 * 10 + i0
      x2 = i2 * 100 + i0 * 10 + i1
      if (is_prime x0) and (is_prime x1) and (is_prime x2)
        list.push x0
        list.push x1
        list.push x2
      i2++
    i1++
  i0++

i0 = 0
while i0 < 10
  i1 = 0
  while i1 < 10
    i2 = 0
    while i2 < 10
      i3 = 0
      while i3 < 10
        x0 = i0 * 1000 + i1 * 100 + i2 * 10 + i3
        x1 = i1 * 1000 + i2 * 100 + i3 * 10 + i0
        x2 = i2 * 1000 + i3 * 100 + i0 * 10 + i1
        x3 = i3 * 1000 + i0 * 100 + i1 * 10 + i2
        if (is_prime x0) and (is_prime x1) and (is_prime x2) and (is_prime x3)
          list.push x0
          list.push x1
          list.push x2
          list.push x3
        i3++
      i2++
    i1++
  i0++

i0 = 0
while i0 < 10
  i1 = 0
  while i1 < 10
    i2 = 0
    while i2 < 10
      i3 = 0
      while i3 < 10
        i4 = 0
        while i4 < 10
          x0 = i0 * 10000 + i1 * 1000 + i2 * 100 + i3 * 10 + i4
          x1 = i1 * 10000 + i2 * 1000 + i3 * 100 + i4 * 10 + i0
          x2 = i2 * 10000 + i3 * 1000 + i4 * 100 + i0 * 10 + i1
          x3 = i3 * 10000 + i4 * 1000 + i0 * 100 + i1 * 10 + i2
          x4 = i4 * 10000 + i0 * 1000 + i1 * 100 + i2 * 10 + i3
          if (is_prime x0) and (is_prime x1) and (is_prime x2) and (is_prime x3) and (is_prime x4)
            list.push x0
            list.push x1
            list.push x2
            list.push x3
            list.push x4
          i4++
        i3++
      i2++
    i1++
  i0++

i0 = 0
while i0 < 10
  i1 = 0
  while i1 < 10
    i2 = 0
    while i2 < 10
      i3 = 0
      while i3 < 10
        i4 = 0
        while i4 < 10
          i5 = 0
          while i5 < 10
            x0 = i0 * 100000 + i1 * 10000 + i2 * 1000 + i3 * 100 + i4 * 10 + i5
            x1 = i1 * 100000 + i2 * 10000 + i3 * 1000 + i4 * 100 + i5 * 10 + i0
            x2 = i2 * 100000 + i3 * 10000 + i4 * 1000 + i5 * 100 + i0 * 10 + i1
            x3 = i3 * 100000 + i4 * 10000 + i5 * 1000 + i0 * 100 + i1 * 10 + i2
            x4 = i4 * 100000 + i5 * 10000 + i0 * 1000 + i1 * 100 + i2 * 10 + i3
            x5 = i5 * 100000 + i0 * 10000 + i1 * 1000 + i2 * 100 + i3 * 10 + i4
            if (is_prime x0) and (is_prime x1) and (is_prime x2) and (is_prime x3) and (is_prime x4) and (is_prime x5)
              list.push x0
              list.push x1
              list.push x2
              list.push x3
              list.push x4
              list.push x5
            i5++
          i4++
        i3++
      i2++
    i1++
  i0++

console.log (array_unique list).length
