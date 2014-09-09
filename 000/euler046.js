/**
 * Goldbach's other conjecture
 *
 * Completed on Tue, 9 Sep 2014, 22:52
 */

(function () {
    var prime_cache = {};

    function is_prime(t) {
        if (t == 1) {
            return false;
        }

        if (t == 2) {
            return true;
        }

        if (prime_cache[t] !== undefined) {
            return prime_cache[t];
        }

        var m = Math.floor(Math.sqrt(t));
        var result = true;

        for (var i = 2; i <= m; i++) {
            if (t % i == 0) {
                result = false;
                break;
            }
        }

        prime_cache[t] = result;

        return result;
    }

    var x = 3;
    var z, goldbach;

    do {
        if (is_prime(x)) {
            goldbach = true;
        } else {
            goldbach = false;

            var max = Math.ceil(Math.sqrt(x * 0.5));
            for (var j = 1; j < max; j++) {
                z = x - 2 * j * j;
                goldbach = goldbach || is_prime(z);
            }
        }

        x += 2;
    } while (goldbach);

    console.log(x - 2);
})();
