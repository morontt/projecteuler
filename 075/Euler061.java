/**
 * Cyclical figurate numbers
 */

import java.util.function.Function;

class Elem {
    int i, lo, hi;

    Elem(int i) {
        this.i = i;

        lo = i % 100;
        hi = (i - lo) / 100;
    }
}

public class Euler061 {
    Elem[][] matrix = new Elem[6][100];
    int[] lenghts = new int[6];

    void fillRow(int idx, Function <Integer, Integer> f) {
        int i = 0, j = 0, r;

        do {
            r = f.apply(i);
            if (r > 999 && r < 10000) {
                Elem e = new Elem(r);
                if (e.lo > 9) {
                    // System.out.println("hi: " + e.hi + " lo: " + e.lo);
                    matrix[idx][j] = e;
                    j++;
                }
            }
            i++;
        } while (r < 10000);

        lenghts[idx] = j;
    }

    public static void main(String[] args) {
        Euler061 eu = new Euler061();

        eu.fillRow(0, n -> n * (n + 1) / 2);
        eu.fillRow(1, n -> n * n);
        eu.fillRow(2, n -> n * (3 * n - 1) / 2);
        eu.fillRow(3, n -> n * (2 * n - 1));
        eu.fillRow(4, n -> n * (5 * n - 3) / 2);
        eu.fillRow(5, n -> n * (3 * n - 2));
    }
}
