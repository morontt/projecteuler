/**
 * Lexicographic permutations
 *
 * Completed on Sun, 21 Sep 2014, 22:29
 */

(function () {
    var n = 10;
    var position = 1000000;

    var list = [],
        used = [],
        count = 0;

    function generate(k) {
        var i;

        if (k == n) {
            count++;
            if (count == position) {
                console.log(list);
            }
        } else {
            for (i = 0; i < n; i++) {
                if (used[i] !== true) {
                    used[i] = true;
                    list[k] = i;

                    generate(k + 1);
                    used[i] = false;
                }
            }
        }
    }

    generate(0);
})();
