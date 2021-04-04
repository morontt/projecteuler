//
// Magic 5-gon ring
//
// Completed on
//

program euler068;

const N = 6;

type
  ring = array[0..(N - 1)] of integer;

var
  r, z : ring;
  i : integer;

procedure show(shape: ring);
var
  g : array[0..2] of integer;
  gr : array[0..2, 0..2] of integer;
  min, idx, k, j : integer;
begin
  g[0] := shape[0];
  g[1] := shape[3];
  g[2] := shape[5];
  gr[0][0] := shape[0];
  gr[0][1] := shape[1];
  gr[0][2] := shape[2];
  gr[1][0] := shape[3];
  gr[1][1] := shape[2];
  gr[1][2] := shape[4];
  gr[2][0] := shape[5];
  gr[2][1] := shape[4];
  gr[2][2] := shape[1];
  min := 32767;
  for i := 0 to 2 do
    begin
      if (g[i] < min) then
        begin
          min := g[i];
          idx := i;
        end;
    end;

  for i := 0 to 2 do
    begin
      k := idx + i;
      if (k > 2) then
        k := k - 3;
      for j := 0 to 2 do
        write(gr[k][j]);
    end;
  writeln;
end;

function check_ring(shape: ring; depth, sum: integer): integer;
var
  t : integer;
begin
  check_ring := sum;
  case (depth) of
    3: check_ring := shape[0] + shape[1] + shape[2];
    5:
      begin
        t := shape[2] + shape[3] + shape[4];
        if (t <> sum) then
          check_ring := -1;
      end;
    6:
      begin
        t := shape[1] + shape[4] + shape[5];
        if (t <> sum) then
          check_ring := -1;
      end;
  end;
end;

procedure permutation(shape, used: ring; depth, sum: integer);
var
  j, s: integer;
begin
  s := check_ring(shape, depth, sum);
  if (s >= 0) then
    begin
      if (depth = N) then
        show(shape)
      else
        begin
          for j := 0 to (N - 1) do
            begin
              if (used[j] = 0) then
                begin
                  used[j] := 1;
                  shape[depth] := j + 1;
                  permutation(shape, used, depth + 1, s);
                  used[j] := 0;
                end;
            end;
        end;
    end;
end;

begin
  for i := 0 to (N - 1) do
    begin
      r[i] := 0;
      z[i] := 0;
    end;

  permutation(r, z, 0, 0);
end.
