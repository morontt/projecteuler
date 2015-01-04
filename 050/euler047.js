/**
 * Distinct primes factors
 *
 * Completed on Sun, 4 Jan 2015, 20:16
 */

(function () {
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

    function div_count(x) {
        var arr = divs(x);
        var i, l = arr.length;

        var used = {};
        var cnt = 0;

        for (i = 0; i < l; i++) {
            if (used[arr[i]] === undefined) {
                used[arr[i]] = true;
                cnt++;
            }
        }

        return cnt;
    }

    var z = 120; // 2 * 3 * 4 * 5
    var loop = true;

    while (loop) {
        z++;
        if (div_count(z) == 4) {
            z++;
            if (div_count(z) == 4) {
                z++;
                if (div_count(z) == 4) {
                    z++;
                    if (div_count(z) == 4) {
                        loop = false;
                        z -= 3;
                    }
                }
            }
        }
    }

    console.log(z);
})();
