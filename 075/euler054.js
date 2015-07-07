/**
 * Poker hands
 */

(function () {
    var fs = require('fs');

    var datalength = 1000;

    var Hand = function Hand() {
        this.suit_c = [];
        this.suit_s = [];
        this.suit_h = [];
        this.suit_d = [];
        this.suits = ['suit_c', 'suit_s', 'suit_h', 'suit_d'];

        var tmp, val;
        for (var i = 0; i < arguments.length; i++) {
            tmp = arguments[i].split('');
            switch (tmp[0]) {
                case 'T':
                    val = 10;
                    break;
                case 'J':
                    val = 11;
                    break;
                case 'Q':
                    val = 12;
                    break;
                case 'K':
                    val = 13;
                    break;
                case 'A':
                    val = 14;
                    break;
                default:
                    val = parseInt(tmp[0]);
            }
            this['suit_' + tmp[1].toLowerCase()].push(val);
        }
    };

    Hand.prototype.in_suit = function (el, suit) {
        return (this[suit].indexOf(el) > -1);
    };

    /**
     * Ten, Jack, Queen, King, Ace, in same suit.
     */
    Hand.prototype.isRoyalFlush = function () {
        var suite;
        for (var i = 0; i < 4; i++) {
            suite = this.suits[i];
            if (this[suite].length == 5
                && this.in_suit(10, suite)
                && this.in_suit(11, suite)
                && this.in_suit(12, suite)
                && this.in_suit(13, suite)
                && this.in_suit(14, suite)
            ) {
                return true;
            }
        }

        return false;
    };

    /**
     * All cards are consecutive values of same suit.
     */
    Hand.prototype.isStraightFlush = function () {
        var max, min, suite;
        for (var i = 0; i < 4; i++) {
            suite = this.suits[i];
            if (this[suite].length == 5) {
                min = this[suite].reduce(function (prev, curr) {
                    return prev > curr ? curr: prev;
                }, 15);

                max = this[suite].reduce(function (prev, curr) {
                    return prev < curr ? curr: prev;
                }, 0);

                if (max == min + 4) {
                    return true;
                }
            }
        }

        return false;
    };

    /**
     * Four cards of the same value.
     */
    Hand.prototype.isFourOfKind = function () {
        var suite, cnt;
        for (var j = 2; j < 15; j++) {
            cnt = 0;
            for (var i = 0; i < 4; i++) {
                suite = this.suits[i];
                if (this.in_suit(j, suite)) {
                    cnt++;
                }
            }

            if (cnt == 4) {
                return true;
            }
        }

        return false;
    };

    function checkHand(string) {
        var cards = string.split(' ');
        var hand1 = new Hand(cards[0], cards[1], cards[2], cards[3], cards[4]);
        var hand2 = new Hand(cards[5], cards[6], cards[7], cards[8], cards[9]);
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
