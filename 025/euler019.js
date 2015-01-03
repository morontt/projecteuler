/**
 * Counting Sundays
 *
 * Completed on Fri, 19 Sep 2014, 02:03
 */

(function () {
    var month = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
    var i, y, delta;

    var t = 0;
    var sum = 0;

    for (i = 0; i < 12; i++) {
        t +=  month[i];
    }
    t = t % 7;

    for (y = 1901; y < 2001; y++) {
        delta = (y % 4 == 0) ? 1 : 0;
        for (i = 0; i < 12; i++) {
            if (t % 7 == 6) {
                sum++;
            }

            if (i == 1) {
                t += month[i] + delta;
            } else {
                t += month[i];
            }
        }
        t = t % 7;
    }

    console.log(sum);
})();
