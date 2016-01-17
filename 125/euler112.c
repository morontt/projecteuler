/*
 * Bouncy numbers
 *
 * Compilation: gcc -Wall euler112.c -o 112
 * Completed on Sun, 6 Dec 2015, 01:36
 */

#include <stdio.h>

int bouncy(long x) {
    int s_old, s = 0;
    long x0, x1;
    int res = 0;

    x0 = x % 10;
    x /= 10;
    x1 = x % 10;
    x /= 10;

    if (x0 > x1) {
        s = 1;
    }

    if (x0 < x1) {
        s = -1;
    }

    s_old = s;

    while (x > 0 && res == 0) {
        x0 = x1;
        x1 = x % 10;
        x /= 10;
        s = 0;

        if (x0 > x1) {
            s = 1;
        }

        if (x0 < x1) {
            s = -1;
        }

        if (s * s_old == -1) {
            res = 1;
        }

        if (s != 0) {
            s_old = s;
        }
    }

    return res;
}

int main() {
    long i = 101, j = 0;
    double r = 0.0;

    while (r < 0.99) {
        if (bouncy(i) == 1) {
            j++;
        }
        r = 1.0 * j / i;
        i++;
    }
    printf("%ld\n", i - 1);

    return 0;
}
