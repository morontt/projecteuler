/**
 * Number letter counts
 *
 * Completed on Thu, 18 Sep 2014, 23:48
 */

(function () {
    var em = [0, 3, 3, 5, 4, 4, 3, 5, 5, 4, 3, 6, 6, 8, 8, 7, 7, 9, 8, 8]; //one, two, three...
    var dm = [0, 3, 6, 6, 5, 5, 5, 7, 6, 6];
    var i, sum = 0;
    var e, d, s;

    for (i = 1; i < 1000; i++) {
        e = i % 10;
        d = ((i % 100) - e) / 10;
        s = (i - 10 * d - e) / 100;

        if (s > 0) {
            sum += em[s] + 7; //hundred
            if (d > 0 || e > 0) {
                sum += 3; //and
            }
        }

        if (d == 0 || d == 1) {
            sum += em[d * 10 + e];
        } else {
            sum += dm[d] + em[e];
        }
    }

    sum += 11; //one thousand

    console.log(sum);
})();
