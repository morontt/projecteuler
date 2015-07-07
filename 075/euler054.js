/**
 * Poker hands
 */

(function () {
    var fs = require('fs');

    var datalength = 10;

    function Hand() {
        this.c = [];
        this.s = [];
        this.h = [];
        this.d = [];

        var tmp;
        for (var i = 0; i < arguments.length; i++) {
            tmp = arguments[i].split('');
            this[tmp[1].toLowerCase()].push(tmp[0]);
        }
    }

    function checkHand(string) {
        var cards = string.split(' ');
        var hand1 = new Hand(cards[0], cards[1], cards[2], cards[3], cards[4]);
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
