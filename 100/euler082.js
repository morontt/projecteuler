/**
 * Path sum: three ways
 *
 * Completed on Sun, 11 Apr 2021, 19:47
 */
(function () {
    var n = 80;

    var fs = require('fs');
    fs.readFile('./p082_matrix.txt', {encoding: 'ascii'}, function (err, data) {
        if (err) {
            throw err;
        }

        var i, j, k, m, arr, s, rows = data.split('\n');
        var matrix = [];
        for (i = 0; i < n; i++) {
            matrix[i] = rows[i].split(',').map(function (e) {
                return parseInt(e)
            });
        }

        // console.table(matrix);

        var w = [];
        for (i = 0; i < n; i++) {
            w.push(new Array(n));
        }

        for (i = 0; i < n; i++) {
            for (j = 0; j < n; j++) {
                w[i][j] = 0;
            }
        }

        for (j = 0; j < n; j++) {
            w[j][n - 1] = matrix[j][n - 1];
        }

        for (m = n - 2; m >= 0; m--) {
            for (i = 0; i < n; i++) {
                arr = [];
                for (j = 0; j < n; j++) {
                    s = 0;
                    if (i === j) {
                        s += w[i][m + 1];
                        s += matrix[i][m];
                    } else if (j < i) {
                        s += w[j][m + 1];
                        for (k = j; k <= i; k++) {
                            s += matrix[k][m];
                        }
                    } else if (j > i) {
                        s += w[j][m + 1];
                        for (k = i; k <= j; k++) {
                            s += matrix[k][m];
                        }
                    }
                    arr.push(s);
                }
                w[i][m] = Math.min.apply(null, arr);
            }
        }

        // console.table(w);

        arr = [];
        for (i = 0; i < n; i++) {
            arr.push(w[i][0]);
        }
        console.log('Answer: ' + Math.min.apply(null, arr));
    });
})();
