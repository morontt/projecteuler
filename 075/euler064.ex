# Odd period square roots
#
# Completed on

defmodule Euler064 do
  def iteration(list, x, i) when x > 0 do
    if i < 20 do
      r = 1.0 / x
      ai = trunc(r)
      xi = r - ai
      iteration [ai | list], xi, i + 1
    else
      list
    end
  end

  def solve() do
    x = :math.sqrt(23.0)
    a0 = trunc(x)
    x0 = x - a0
    l = iteration [], x0, 0
    IO.inspect(l)
  end
end
