<?php

$fp = fopen('p096_sudoku.txt', 'r');
$idx = 0;
$out = null;
while (($buffer = fgets($fp, 4096)) !== false) {
    if (strpos($buffer, 'Grid') !== false) {
        if ($out !== null) {
            fclose($out);
        }
        $idx++;
        $out = fopen('./variants/grid_' . $idx . '.txt', 'a');
    }

    echo $buffer;
    fwrite($out, $buffer);
}
fclose($fp);
fclose($out);
