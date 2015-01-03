/**
 * Even Fibonacci numbers
 *
 * Completed on Mon, 8 Sep 2014, 23:22
 */

var f_a = 1,
    f_b = 1,
    f_c = 0,
    f_summ = 0,
    max = 4000000;

do {
    if (f_c % 2 == 0) {
        f_summ += f_c;
    }

    f_c = f_a + f_b;
    f_a = f_b;
    f_b = f_c;
} while (f_c < max);

console.log(f_summ);
