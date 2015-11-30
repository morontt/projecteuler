/*
 * Ordered fractions
 *
 * Compilation: gcc -Wall euler071.c -o 71 -lm
 * Completed on Mon, 30 Nov 2015, 23:36
 */

#include <stdio.h>
#include <math.h>

int main () {
    int al = 2, bl = 5;
    int ah = 3, bh = 7;
    int i = 9;
    int j, z, x;

    double rl = 1.0 * al / bl;
    double rh = 1.0 * ah / bh;
    double r;

    while (i < 1000001) {
        x = (int) floor(i * rl);
        z = (int) ceil(i * rh);

        for (j = x; j < z; j++) {
            r = 1.0 * j / i;
            if ((r > rl) && (r < rh)) {
                al = j;
                bl = i;
                rl = 1.0 * al / bl;
            }
        }

        i++;
    }

    printf("%d %d\n", al, bl);
    return 0;
}
