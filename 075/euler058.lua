-- Spiral primes
--
-- Completed on Fri, 31 Jul 2015, 22:10

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

prime = 0
count = 1
ratio = 1.0

n = 3
while ratio > 0.1 do
  local d = n * n
  for j = 1, 3 do
    d = d - n + 1
    if is_prime(d) then
      prime = prime + 1
    end
  end
  count = count + 4
  ratio = prime / count
  n = n + 2
end

print(n - 2)
