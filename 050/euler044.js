/**
 * Pentagon numbers
 *
 * Completed on Tue, 6 Jan 2015, 01:04
 */

(function () {
    var limit = 10000;

    var i, j, p, p1, p2;
    var pentagon_cache = {};
    var pentagons = [];

    for (i = 1; i < limit; i++) {
        p = i * (3 * i - 1) / 2;
        pentagons.push(p);
        pentagon_cache[p] = true;
    }

    var s1 = Math.floor(Math.sqrt(limit * limit / 2));
    var s0 = s1 - 1;

    for (i = 1; i < s0; i++) {
        for (j = i + 1; j < s1; j++) {
            p1 = pentagons[j] - pentagons[i];
            p2 = pentagons[j] + pentagons[i];
            if (pentagon_cache[p2] !== undefined && pentagon_cache[p1] !== undefined) {
                console.log(p1);
                process.exit();
            }
        }
    }
})();
