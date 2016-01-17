/*
 * Counting rectangles
 *
 * Compilation: gcc -Wall euler085.c -o 85
 * Completed on Mon, 7 Dec 2015, 23:40
 */

#include <stdio.h>

#define TARGET 2000000

unsigned long rectangles(long a, long b) {
    unsigned long s = 0;
    long a1, b1;
    long i, j;

    a1 = a + 1;
    b1 = b + 1;

    for (i = 1; i < a1; i++) {
        for (j = 1; j < b1; j++) {
            s += (a1 - i) * (b1 - j);
        }
    }

    return s;
}

int main() {
    long i, j, s;
    unsigned long r = 0, r0 = 0, d, dmin = 1000;

    i = 1;
    do {
        j = 1;
        r = 0;
        while (r < TARGET) {
            r0 = r;
            r = rectangles(j, i);
            j++;
        }

        d = (TARGET - r) * (TARGET - r);
        if (d < dmin) {
            dmin = d;
            s = (j - 1) * i;
        }
        d = (TARGET - r0) * (TARGET - r0);
        if (d < dmin) {
            dmin = d;
            s = (j - 2) * i;
        }

        i++;
    } while (j >= i);

    printf("%ld\n", s);
    return 0;
}
