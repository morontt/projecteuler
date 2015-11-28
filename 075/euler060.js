/**
 * Prime pair sets
 *
 * Completed on Sun, 29 Nov 2015, 00:07
 */
(function () {
    var nmax = 1100;
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

    function check_prime_pair(a, b) {
        return is_prime(parseInt(a.toString() + b.toString())) && is_prime(parseInt(b.toString() + a.toString()));
    }

    var primes_dict = [2];
    var i = 3;

    while (primes_dict.length < nmax) {
        if (is_prime(i)) {
            primes_dict.push(i);
        }
        i += 2;
    }

    function find_pairs(m) {
        var d = [];
        var i1, i2;
        var m1 = m - 1;
        for (i1 = 0; i1 < m1; i1++) {
            for (i2 = i1 + 1; i2 < m; i2++) {
                if (check_prime_pair(primes_dict[i1], primes_dict[i2])) {
                    d.push([i1, i2]);
                }
            }
        }

        return d;
    }

    function find_triples(m) {
        var pairs = find_pairs(m - 1);
        var idx, i, l = pairs.length;
        var d = [];

        for (idx = 0; idx < l; idx++) {
            for (i = pairs[idx][1] + 1; i < m; i++) {
                if (check_prime_pair(primes_dict[i], primes_dict[pairs[idx][0]])
                    && check_prime_pair(primes_dict[i], primes_dict[pairs[idx][1]])
                ) {
                    d.push([pairs[idx][0], pairs[idx][1], i]);
                }
            }
        }

        return d;
    }

    function find_fours(m) {
        var triples = find_triples(m - 1);
        var idx, i, l = triples.length;
        var d = [];

        for (idx = 0; idx < l; idx++) {
            for (i = triples[idx][2] + 1; i < m; i++) {
                if (check_prime_pair(primes_dict[i], primes_dict[triples[idx][0]])
                    && check_prime_pair(primes_dict[i], primes_dict[triples[idx][1]])
                    && check_prime_pair(primes_dict[i], primes_dict[triples[idx][2]])
                ) {
                    d.push([triples[idx][0], triples[idx][1], triples[idx][2], i]);
                }
            }
        }

        return d;
    }

    var fours = find_fours(nmax - 1);

    var idx, x, l = fours.length;
    for (idx = 0; idx < l; idx++) {
        for (i = fours[idx][3] + 1; i < nmax; i++) {
            x = primes_dict[i];

            if (check_prime_pair(x, primes_dict[fours[idx][3]])
                && check_prime_pair(x, primes_dict[fours[idx][2]])
                && check_prime_pair(x, primes_dict[fours[idx][1]])
                && check_prime_pair(x, primes_dict[fours[idx][0]])
            ) {
                console.log(
                    primes_dict[fours[idx][0]],
                    primes_dict[fours[idx][1]],
                    primes_dict[fours[idx][2]],
                    primes_dict[fours[idx][3]],
                    x
                );
                console.log(
                    primes_dict[fours[idx][0]]
                    + primes_dict[fours[idx][1]]
                    + primes_dict[fours[idx][2]]
                    + primes_dict[fours[idx][3]]
                    + x
                );
            }
        }
    }
})();
