# Odd period square roots
#
# Completed on

defmodule Euler064 do
  def iteration(list, x, eps) when eps < 0.01 and x > 0 do
    r = 1.0 / x
    ai = trunc(r)
    xi = r - ai
    iteration [ai | list], xi, eps / (x * x)
  end

  def iteration(list, _x, _eps) do
    Enum.reverse(list)
  end

  def check(n) do
    x = :math.sqrt(1.0 * n)
    a0 = trunc(x)
    x0 = x - a0
    l = iteration [], x0, 1.0e-15
    IO.inspect(n)
    IO.inspect(l)
  end

  def solve do
    limit = 10
    ch1 = for val <- 2..(limit * limit), do: val
    ch2 = for val <- 1..limit, do: val * val
    ch3 = ch1 -- ch2
    Enum.each(ch3, &check/1)
  end
end
