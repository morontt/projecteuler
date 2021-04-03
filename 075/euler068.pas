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
begin
  for i := 0 to (N - 1) do
    write(shape[i]);
  writeln('');
end;

procedure permutation(shape, used: ring; depth: integer);
var
  j: integer;
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
              permutation(shape, used, depth + 1);
              used[j] := 0;
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

  permutation(r, z, 0);
end.
