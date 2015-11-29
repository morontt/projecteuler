/**
 * Path sum: two ways
 *
 * Completed on Sun, 29 Nov 2015, 12:02
 */
(function () {
    var n = 80;

    var fs = require('fs');
    fs.readFile('./p081_matrix.txt', {encoding: 'ascii'}, function (err, data) {
        if (err) {
            throw err;
        }

        var i, j, rows = data.split('\n');
        var matrix = [];
        for (i = 0; i < n; i++) {
            matrix[i] = rows[i].split(',').map(function (e) {
                return parseInt(e)
            });
        }

        var w = [];
        for (i = 0; i < n; i++) {
            w.push(new Array(n));
        }

        for (i = 0; i < n; i++) {
            for (j = 0; j < n; j++) {
                w[i][j] = 0;
            }
        }

        w[n - 1][n - 1] = matrix[n - 1][n - 1];

        for (i = n - 2; i >= 0; i--) {
            w[n - 1][i] = matrix[n - 1][i] + w[n - 1][i + 1];
            w[i][n - 1] = matrix[i][n - 1] + w[i + 1][n - 1];
        }

        for (i = n - 2; i >= 0; i--) {
            for (j = n - 2; j >= 0; j--) {
                w[i][j] = matrix[i][j] + Math.min(w[i][j + 1], w[i + 1][j]);
            }
        }

        console.log(w[0][0]);
    });
})();
