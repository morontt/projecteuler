# Odd period square roots
#
# Completed on

defmodule Euler064 do
  def iteration(list, x, eps) when eps < 0.1 and x > 0 do
    r = 1.0 / x
    ai = trunc(r)
    xi = r - ai
    iteration [ai | list], xi, eps / (x * x)
  end

  def iteration(list, _x, _eps) do
    list
  end

  def solve() do
    x = :math.sqrt(23.0)
    a0 = trunc(x)
    x0 = x - a0
    l = iteration [], x0, 1.0e-15
    IO.inspect(l)
  end
end
