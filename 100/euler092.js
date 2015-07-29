// completed 29.07.2015 13:45:34
// 8581146

(function () {
    var cache = {};

    function sum_sq(x) {
        var res, s, t, q = 0;

        if (cache[x] !== undefined) {
            return cache[x];
        }

        if (x == 1) {
            res = 0;
        } else {
            if (x == 89) {
                res = 1;
            } else {
                s = x.toString().split('');
                s.forEach(function (el) {
                    t = parseInt(el);
                    q += t * t;
                });
                res = sum_sq(q);
            }
        }

        cache[x] = res;

        return res;
    }

    var sum = 0, i = 1;
    
    while (i < 10000000) {
        sum += sum_sq(i);
        i++;
    }

    console.log(sum);
})();
