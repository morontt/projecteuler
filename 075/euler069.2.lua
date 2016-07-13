-- Totient maximum
--
-- Completed on Sat, 14 Jul 2016, 01:23

function is_prime (x)
  local res = true
  local lim = math.floor(math.sqrt(x))
  if x % 2 == 0 then
    return false
  end
  for i = 3, lim, 2 do
    if x % i == 0 then
      res = false
      break
    end
  end
  return res
end

s = 6
old = 0
n = 3
while s < 1000000 do
  n = n + 1
  if is_prime(n) then
    old = s
    s = s * n
  end
end

print(old)
