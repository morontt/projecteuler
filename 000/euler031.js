/**
 * Coin sums
 *
 * Completed on Wed, 17 Sep 2014, 21:04
 */

(function () {
    var n = 200;
    var list = [200, 100, 50, 20, 10, 5, 2, 1];

    function comb(k, s) {
        if (k == 0) {
            return 1;
        }

        if (s.length == 1) {
            return (k % s[0] == 0) ? 1 : 0;
        }

        var i, d, rest, u, sum = 0;

        d = Math.floor(k/s[0]) + 1;
        u = s.slice(1);

        for (i = 0; i < d; i++) {
            rest = k - s[0] * i;
            sum += comb(rest, u);
        }

        return sum;
    }

    console.log(comb(n, list));
})();
