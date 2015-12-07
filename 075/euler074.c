/*
 * Digit factorial chains
 *
 * Compilation: gcc -Wall euler074.c -o 74
 * Completed on Sat, 5 Dec 2015, 14:45
 */

 #include <stdio.h>

long f[] = {1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880};

long factsum(long x) {
    long s = 0;

    while (x > 0) {
        s += f[x % 10];
        x /= 10;
    }

    return s;
}

int main() {
    long i, t;
    int j, k, find, rez;
    long chain[70];

    rez = 0;

    for (i = 1; i < 1000001; i++) {
        find = 0;
        j = 1;
        chain[0] = i;
        t = i;
        while (find == 0) {
            t = factsum(t);
            for (k = 0; k < j; k++) {
                if (t == chain[k]) {
                    find = j;
                    break;
                }
            }
            chain[j] = t;
            j++;
        }
        if (find == 60) {
            rez++;
        }
    }

    printf("%d\n", rez);
    return 0;
}
