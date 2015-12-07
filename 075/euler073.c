/*
 * Counting fractions in a range
 *
 * Compilation: gcc -Wall euler073.c -o 73
 * Completed on Sat, 5 Dec 2015, 22:27
 */

#include <stdio.h>

#define HIGHT 12001

int hcf(int a, int b) {
    int r;

    r = a % b;
    while (r > 0) {
        a = b;
        b = r;
        r = a % b;
    }

    return b;
}

int main() {
    int i = 3;
    long s = 0;
    int j, z, x;

    while (i < HIGHT) {
        x = 1 + i / 3;
        if (i % 2 == 1) {
            z = 1 + i / 2;
        } else {
            z = i / 2;
        }

        for (j = x; j < z; j++) {
            if (hcf(i, j) == 1) {
                s += 1;
            }
        }

        i++;
    }

    printf("%ld\n", s);
    return 0;
}
