/**
 * Number spiral diagonals
 *
 * Completed on Tue, 23 Sep 2014, 01:15
 */

(function () {
    var n = 1001;

    var s = 1;
    var x = 1;
    var d = 1;
    var nkv = n * n;

    do {
        s += 4 * x + 20 * d;
        x += 8 * d;
        d++;
    } while (x < nkv);

    console.log(s);
})();
