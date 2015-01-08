/**
 * Prime digit replacements
 *
 * Completed on Fri, 9 Jan 2015, 00:26
 */

(function () {
    var prime_cache = {};
    var max_cnt = 0;

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

    function test_asterisk(arr) {
        var str = arr.join('');
        var i, t, cnt = 0;

        for (i = 0; i < 10; i++) {
            t = parseInt(str.replace(/\*/g, i.toString()));
            if (is_prime(t) && !(i == 0 && arr[0] == '*')) {
                cnt++;
            }
        }

        if (cnt >= max_cnt) {
            max_cnt = cnt;
            console.log(cnt);
            console.log(str);
        }
    }

    function replace_comb_1(x) {
        var str = x.toString();
        var arr = str.split('');
        var len = str.length;
        var i, tmp;

        for (i = 0; i < len; i++) {
            tmp = arr.slice(0);
            tmp.splice(i, 0, '*');
            test_asterisk(tmp);
        }
    }

    function replace_comb_2(x) {
        var str = x.toString();
        var arr = str.split('');
        var len = str.length;
        var i, j, tmp;

        for (i = 0; i < len; i++) {
            for (j = i; j < len; j++) {
                tmp = arr.slice(0);
                tmp.splice(i, 0, '*');
                tmp.splice(j + 1, 0, '*');
                test_asterisk(tmp);
            }
        }
    }

    function replace_comb_3(x) {
        var str = x.toString();
        var arr = str.split('');
        var len = str.length;
        var i, j, k, tmp;

        for (i = 0; i < len; i++) {
            for (j = i; j < len; j++) {
                for (k = j; k < len; k++) {
                    tmp = arr.slice(0);
                    tmp.splice(i, 0, '*');
                    tmp.splice(j + 1, 0, '*');
                    tmp.splice(k + 2, 0, '*');
                    test_asterisk(tmp);
                }
            }
        }
    }

    var limit = 500;
    var n = 3;

    while (n < limit) {
        replace_comb_1(n);
        replace_comb_2(n);
        replace_comb_3(n);
        n += 2;
    }
})();
