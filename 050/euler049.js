/**
 * Prime permutations
 *
 * Completed on Mon, 5 Jan 2015, 00:06
 */

(function () {
    var prime_cache = {};
    var permutation_cache = {};

    function is_prime(t) {
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

    function permutation_hash(t) {
        if (permutation_cache[t] !== undefined) {
            return permutation_cache[t];
        }

        var st = t.toString().split('').sort().join('');

        permutation_cache[t] = st;

        return st;
    }

    var primes = [];
    var hashes = [];

    var i = 1001;
    while (i < 10000) {
        if (is_prime(i)) {
            primes.push(i);
            hashes.push(permutation_hash(i))
        }
        i++;
    }

    var len = primes.length;
    var len1 = len - 1;
    var j, d, q;

    for (i = 0; i < len1; i++) {
        for (j = i + 1; j < len; j++) {
            if (hashes[i] === hashes[j]) {
                d = primes[j] - primes[i];
                q = primes[j] + d;
                if (hashes[j] === permutation_hash(q) && is_prime(q)) {
                    console.log(primes[i] + '-' + primes[j] + '-' + q);
                }
            }
        }
    }
})();
