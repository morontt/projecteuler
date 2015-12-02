/**
 * Counting summations
 *
 * Completed on Thu, 3 Dec 2015, 00:38
 */

(function () {
    var n = 100;
    var list = [];
    var j;

    for (j = 99; j > 0; j--) {
        list.push(j);
    }

    function comb(k, s) {
        if (k == 0) {
            return 1;
        }

        if (s.length == 1) {
            return (k % s[0] == 0) ? 1 : 0;
        }

        var i, d, rest, u, sum = 0;

        d = Math.floor(k / s[0]) + 1;
        u = s.slice(1);

        for (i = 0; i < d; i++) {
            rest = k - s[0] * i;
            sum += comb(rest, u);
        }

        return sum;
    }

    console.log(comb(n, list));
})();
