/**
 * Highly divisible triangular number
 *
 * Completed on Fri, 12 Sep 2014, 14:41
 */

(function () {
    var d = 500;

    function div(x) {
        var i, limit = Math.ceil(Math.sqrt(x));

        for (i = 2; i < limit; i++) {
            if (x % i == 0) {
                return i;
            }
        }

        return false;
    }

    function divs(x) {
        var divs = [];
        var r;

        do {
            r = div(x);
            if (r !== false) {
                x = x / r;
                divs.push(r);
            } else {
                divs.push(x);
            }
        } while (r);

        return divs;
    }

    function combinations(x) {
        var arr = divs(x);
        var l = arr.length;
        var m = Math.pow(2, l) - 1;
        var i, j, p, tmp;

        var c = 2;
        tmp = {};

        for (i = 1; i < m; i++) {
            p = 1;
            for (j = 0; j < l; j++) {
                if ((i >> j) % 2 == 1) {
                    p *= arr[j];
                }
            }

            if (tmp[p] === undefined) {
                tmp[p] = true;
                c++;
            }
        }

        return c;
    }

    var n = 1;
    var t, cnt;

    do {
        t = n * (n + 1) / 2;
        cnt = combinations(t);

        n++;
    } while (d >= cnt);

    console.log(t);
})();
