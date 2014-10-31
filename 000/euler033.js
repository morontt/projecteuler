/**
 * Digit cancelling fractions
 *
 * Completed on Fri, 31 Oct 2014, 20:32
 */

(function () {
    var i, j, k, a, b;
    var aa = [], bb = [];
    var aaa = 1, bbb = 1;

    function check(x, y, k0, j0) {
        if (x < y && (x * k0 == y * j0)) {
            aa.push(x);
            bb.push(y);
        }
    }

    for (i = 1; i < 10; i++) {
        for (j = 1; j < 10; j++) {
            for (k = 1; k < 10; k++) {
                a = i * 10 + j;
                b = i * 10 + k;
                check(a, b, k, j);

                a = i * 10 + j;
                b = k * 10 + i;
                check(a, b, k, j);

                a = j * 10 + i;
                b = k * 10 + i;
                check(a, b, k, j);

                a = j * 10 + i;
                b = i * 10 + k;
                check(a, b, k, j);
            }
        }
    }

    aa.forEach(function (val, idx) {
        aaa *= val;
        bbb *= bb[idx];
    });

    console.log(aaa);
    console.log(bbb);
})();
