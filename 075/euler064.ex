# Odd period square roots
#
# Completed on Sun, 30 Apr 2017, 19:45

defmodule Euler064 do
  def iteration(_x, _s, start, a0, a1, b0, b1, n) when n > 1 and start == b0 and a0 == 1 and a1 == 1 and b1 == 1 do
    n - 1
  end

  def iteration(x, s, start, a0, a1, b0, b1, n) do
    t0 = 1.0 / (a0 * s / a1 - b0 * 1.0 / b1) |> trunc
    a10 = a1 * a0 * b1 * b1
    a11 = a0 * a0 * b1 * b1 * x - a1 * a1 * b0 * b0
    b10 = t0 * a11 - a1 * a1 * b0 * b1
    {a20, a21} = reduct(a10, a11)
    {b20, b21} = reduct(b10, a11)
    iteration x, s, start, a20, a21, b20, b21, n + 1
  end

  def gcd(m, 0) do
    m
  end

  def gcd(m, n) do
    gcd(n, rem(m, n))
  end

  def reduct(m, n) do
    g = if m > n do
      gcd(m, n)
    else
      gcd(n, m)
    end

    {div(m, g), div(n, g)}
  end

  def check(x) do
    s = :math.sqrt(1.0 * x)
    b = trunc(s)
    iteration x, s, b, 1, 1, b, 1, 1
  end

  def solve do
    limit = 100
    ch1 = for val <- 2..(limit * limit), do: val
    ch2 = for val <- 1..limit, do: val * val
    ch3 = ch1 -- ch2
    Enum.map(ch3, &check/1)
    |> Enum.filter(&(rem(&1, 2) == 1))
    |> length
    |> IO.puts
  end
end
