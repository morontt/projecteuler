//
// Magic 5-gon ring
//
// Completed on Sun, 4 Apr 2021, 10:29
//

program euler068;

const
  N = 10;
  ED = 4;

type
  ring = array[0..(N - 1)] of integer;

var
  r, z : ring;
  i : integer;

procedure show(shape: ring);
var
  g : array[0..ED] of integer;
  gr : array[0..ED, 0..2] of integer;
  min, idx, k, j : integer;
begin
  g[0] := shape[0];
  g[1] := shape[3];
  g[2] := shape[5];
  g[3] := shape[7];
  g[4] := shape[9];
  gr[0][0] := shape[0];
  gr[0][1] := shape[1];
  gr[0][2] := shape[2];
  gr[1][0] := shape[3];
  gr[1][1] := shape[2];
  gr[1][2] := shape[4];
  gr[2][0] := shape[5];
  gr[2][1] := shape[4];
  gr[2][2] := shape[6];
  gr[3][0] := shape[7];
  gr[3][1] := shape[6];
  gr[3][2] := shape[8];
  gr[4][0] := shape[9];
  gr[4][1] := shape[8];
  gr[4][2] := shape[1];
  min := 32767;
  for i := 0 to ED do
    begin
      if (g[i] < min) then
        begin
          min := g[i];
          idx := i;
        end;
    end;

  for i := 0 to ED do
    begin
      k := idx + i;
      if (k > ED) then
        k := k - (ED + 1);
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
    2:
      if (shape[1] = 10) then
        check_ring := -1;
    3:
      begin
        if (shape[2] = 10) then
          check_ring := -1
        else
          check_ring := shape[0] + shape[1] + shape[2];
      end;
    5:
      begin
        if (shape[4] = 10) then
          check_ring := -1
        else
          begin
            t := shape[2] + shape[3] + shape[4];
            if (t <> sum) then
              check_ring := -1;
          end;
      end;
    7:
      begin
        if (shape[6] = 10) then
          check_ring := -1
        else
          begin
            t := shape[4] + shape[5] + shape[6];
            if (t <> sum) then
              check_ring := -1;
          end;
      end;
    9:
      begin
        if (shape[8] = 10) then
          check_ring := -1
        else
          begin
            t := shape[6] + shape[7] + shape[8];
            if (t <> sum) then
              check_ring := -1;
          end;
      end;
    10:
      begin
        t := shape[8] + shape[9] + shape[1];
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
