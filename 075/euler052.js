/**
 * Permuted multiples
 *
 * Completed on Sat, 3 Jan 2015, 19:45
 */

(function () {
    function number_hash(n) {
        var arr = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
        var rest = n;
        var tmp, hash;

        while (rest > 0) {
            tmp = rest % 10;
            arr[tmp] = 1;
            rest = (rest - tmp) / 10;
        }

        hash = arr[0] + 2 * arr[1] + 4 * arr[2] + 8 * arr[3] + 16 * arr[4] + 32 * arr[5] + 64 * arr[6] + 128 * arr[7];
        hash += 256 * arr[8] + 512 * arr[9];

        return hash;
    }

    function check_multiples(x) {
        var h1 = number_hash(x);

        if (h1 != number_hash(2 * x)) {
            return false;
        }

        if (h1 != number_hash(3 * x)) {
            return false;
        }

        if (h1 != number_hash(4 * x)) {
            return false;
        }

        if (h1 != number_hash(5 * x)) {
            return false;
        }

        return h1 == number_hash(6 * x);
    }

    var i = 1;
    var test = true;
    while (test) {
        i++;
        test = !check_multiples(i);
    }

    console.log(i);
})();
