/**
 * Special Pythagorean triplet
 *
 * Created by morontt on 09.09.14.
 */

(function () {
    function check(a, b) {
        return ((500000 + a * b) == 1000 * (a + b));
    }

    function result(x1, m1) {
        var x2, x3;

        x2 = m1 - x1;
        x3 = 1000 - m1;

        console.log(x1*x2*x3);
    }

    var m, x, mp;

top:
    for (m = 500; m < 1000; m++) {
        mp = m / 2;
        for (x = 1; x < m; x++) {
            if (check(x, m - x)) {
                result(x, m);
                break top;
            }
        }
    };
})();
