-- Totient maximum
--
-- incomplete

function gcd (m, n)
  local a = m
  local b = n
  while b > 0 do
    a, b = b, a % b
  end
  return a
end

max = 0
idx = 0

for i = 2, 30000 do
  local count = 0
  local lim = i - 1
  for j = 1, lim do
    if gcd(i, j) == 1 then
      count = count + 1
    end
  end
  -- print(i .. ": " .. count)
  local ratio = i / count
  if ratio > max then
    max = ratio
    idx = i
  end
end

print(idx)
