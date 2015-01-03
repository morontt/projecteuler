/**
 * Combinatoric selections
 *
 * Completed on Sun, 4 Jan 2015, 00:10
 */

(function () {
    var i, j, t;
    var n = 101;
    var c = [];
    var s = 0;

    for (i = 0; i <= n; i++) {
        c[i] = [];
        c[i][0] = 0;
        c[i][i + 1] = 0;
    }

    c[1][1] = 1;

    for (i = 2; i <= n; i++) {
        for (j = 1; j <= i; j++) {
            t = c[i - 1][j - 1] + c[i - 1][j];
            c[i][j] = t;
            if (t > 1000000) {
                s++;
            }
        }
    }

    console.log(s);
})();
