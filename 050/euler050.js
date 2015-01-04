/**
 * Consecutive prime sum
 *
 * Completed on Sun, 4 Jan 2015, 11:58
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

    var max_sum = 1000000;
    var limit = 48000; //about 1000000 / 21
    var d = 21;
    var primes = [2];
    var i = 3;
    var s = 2;

    while (i < limit) {
        if (is_prime(i)) {
            s += i;
            primes.push(s);
        }
        i += 2;
    }

    var max = 0;
    var max_length = 0;
    var prime_length = primes.length;
    var j, i0, z, l;

    i = 0;
    i0 = prime_length - d + 1;
    while (i < i0) {
        j = i + d;
        while (j < prime_length) {
            z = primes[j] - primes[i];
            if (z < max_sum && is_prime(z)) {
                l = j - i;
                if (l >= max_length) {
                    max_length = l;
                    max = z;
                }
            }
            j++;
        }
        i++;
    }

    console.log(max);
    console.log(max_length);
})();
