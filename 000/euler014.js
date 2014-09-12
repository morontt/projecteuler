/**
 * Longest Collatz sequence
 *
 * Completed on Fri, 12 Sep 2014, 15:18
 */

(function () {
    var cache = {};

    function collatz(x) {
        if (x == 1) {
            return 1;
        }

        if (cache[x] !== undefined) {
            return cache[x];
        }

        var next, res;

        if (x % 2 == 0) {
            next = x / 2;
        } else {
            next = 3 * x + 1;
        }

        res = 1 + collatz(next);
        cache[x] = res;

        return res;
    }

    var i, l, big, max = 0;

    for (i = 2; i < 1000000; i++) {
        l = collatz(i);
        if (l > max) {
            max = l;
            big = i;
        }
    }

    console.log(big);
})();
