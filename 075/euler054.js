/**
 * Poker hands
 */

(function () {
    var fs = require('fs');

    var datalength = 1000;

    function checkHand(string) {
        var cards = string.split(' ');
    }

    fs.readFile('./p054_poker.txt', {encoding: 'ascii'}, function (err, data) {
        if (err) {
            throw err;
        }

        var str;
        for (var i = 0; i < datalength; i++) {
            str = '';
            for (var j = 0; j < 30; j++) {
                if (j < 29) {
                    str += data[30 * i + j];
                }
            }
            checkHand(str);
        }
    });
})();
