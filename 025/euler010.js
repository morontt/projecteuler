/**
 * Summation of primes
 *
 * Completed on Tue, 9 Sep 2014, 23:27
 */

(function () {
    var n = 2000000;

    function is_prime(t) {
        var m = Math.floor(Math.sqrt(t));
        var result = true;

        for (var i = 2; i <= m; i++) {
            if (t % i == 0) {
                result = false;
                break;
            }
        }

        return result;
    }

    var x = 3;
    var summ = 2;

    do {
        if (is_prime(x)) {
            summ += x;
        }
        x += 2;
    } while (x < n);

    console.log(summ);
})();
