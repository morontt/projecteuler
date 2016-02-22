/**
 * Cyclical figurate numbers
 *
 * Completed on Sat, 20 Feb 2016, 11:53
 */

import java.util.function.Function;

public class Euler061 {
    Elem[][] matrix = new Elem[6][100];
    int[] lenghts = new int[6];
    boolean[] free = new boolean[]{true, true, true, true, true};
    boolean solved = false;
    
    class Elem {
        int i, lo, hi;

        Elem(int i) {
            this.i = i;

            lo = i % 100;
            hi = (i - lo) / 100;
        }
    }

    void fillRow(int idx, Function <Integer, Integer> f) {
        int i = 0, j = 0, r;

        do {
            r = f.apply(i);
            if (r > 999 && r < 10000) {
                Elem e = new Elem(r);
                if (e.lo > 9) {
                    matrix[idx][j] = e;
                    j++;
                }
            }
            i++;
        } while (r < 10000);

        lenghts[idx] = j;
    }

    void showResult(String str) {
        String[] strs = str.split(" ");
        int result = 0;
        for (String substr : strs) {
            result += Integer.parseInt(substr);
        }
        System.out.println(str);
        System.out.println(result);
    }

    void iteration(int cnt, Elem curr, Elem first, String str) {
        Elem z;
        int next = cnt + 1;
        for (int i = 0; i < 5; i++) {
            if (free[i]) {
                free[i] = false;
                for (int j = 0; j < lenghts[i]; j++) {
                    z = matrix[i][j];
                    if (cnt < 4) {
                        if (z.hi == curr.lo) {
                            iteration(next, z, first, str + z.i + " ");
                        }
                    } else {
                        if (z.hi == curr.lo && z.lo == first.hi) {
                            str += z.i + " " + first.i;
                            showResult(str);
                            solved = true;
                        }
                    }
                }
                free[i] = true;
            }
        }
    }

    public static void main(String[] args) {
        Euler061 eu = new Euler061();

        eu.fillRow(0, n -> n * (n + 1) / 2);
        eu.fillRow(1, n -> n * n);
        eu.fillRow(2, n -> n * (3 * n - 1) / 2);
        eu.fillRow(3, n -> n * (2 * n - 1));
        eu.fillRow(4, n -> n * (5 * n - 3) / 2);
        eu.fillRow(5, n -> n * (3 * n - 2));

        for (int j = 0; j < eu.lenghts[5]; j++) {
            eu.iteration(0, eu.matrix[5][j], eu.matrix[5][j], "");
            if (eu.solved) {
                break;
            }
        }
    }
}
