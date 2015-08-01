-- Totient maximum
--
-- Completed on Sat, 1 Aug 2015, 15:06

function gcd (m, n)
  local a = m
  local b = n
  while b > 0 do
    a, b = b, a % b
  end
  return a
end

function divisors (x)
  local d = {}
  local flag = true
  local m = x
  while flag do
    flag = false
    local z = m
    local lim = math.floor(math.sqrt(m))
    for i = 2, lim do
      local r = z % i
      if r == 0 then
        z = i
        flag = true
        break
      end
    end
    table.insert(d, z)
    m = m / z
  end
  return d
end

function array_unique (t)
  local set = {}
  local res = {}
  for _, k in ipairs(t) do
    set[k] = 1
  end
  for w, _ in pairs(set) do
    table.insert(res, w)
  end
  return res
end

function weigth (t)
  local s = 0.0
  for i = 1, #t do
    s = s + 1.0 / t[i]
  end
  return s
end

function compare (a, b)
  return a['w'] > b['w']
end

wset = {}

for i = 2, 1000000 do
  local divs = array_unique(divisors(i))
  local w = weigth(divs)
  table.insert(wset, {el = i, w = w})
end

table.sort(wset, compare)

max = 0
idx = 0

for i = 1, 100 do
  local u = wset[i]['el']
  local count = 0
  local lim = u - 1
  for j = 1, lim do
    if gcd(u, j) == 1 then
      count = count + 1
    end
  end
  local ratio = u / count
  if ratio > max then
    max = ratio
    idx = u
  end
end

print(idx)
