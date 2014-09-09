/**
 * 10001st prime
 *
 * Completed on Tue, 9 Sep 2014, 22:12
 */

(function () {
    var n = 10001;

    var x = 3;
    var k = 2;

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

    while (k < n) {
        x += 2;
        if (is_prime(x)) {
            k++;
        }
    }

    console.log(x);
})();
