/**
 * Pandigital prime
 *
 * Completed on Sat, 3 Jan 2015, 15:35
 */

(function () {
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

    function check_prime(arr) {
        var i, len = arr.length;
        var val = 0;

        for (i = 0; i < len; i++) {
            val *= 10;
            val += arr[i];
        }

        if (is_prime(val) && max_prime < val) {
            max_prime = val;
        }
    }

    function permutation(inarray, outarray) {
        var i;
        var len = inarray.length;
        var tmp1, tmp2;

        if (len == 0) {
            check_prime(outarray);
        } else {
            for (i = 0; i < len; i++) {
                tmp1 = inarray.slice(0);
                tmp2 = outarray.slice(0);

                tmp2.push(tmp1[i]);
                tmp1.splice(i, 1);

                permutation(tmp1, tmp2);
            }
        }
    }

    /**
     * the sum 1 + 2 + ... + 8 and 1 + 2 + ... + 9 multiple of 3. not prime
     */
    var list = [1, 2, 3, 4, 5, 6, 7];
    var max_prime = 0;

    permutation(list, []);

    console.log(max_prime);
})();
