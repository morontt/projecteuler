/**
 * Poker hands
 *
 * Completed on Thu, 9 Jul 2015, 01:25
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

    Hand.prototype.scoreSameValue = function (z) {
        var suite, cnt, exclude = 0;

        if (arguments.length == 2) {
            exclude = arguments[1];
        }

        for (var j = 14; j > 1; j--) {
            if (j != exclude) {
                cnt = 0;
                for (var i = 0; i < 4; i++) {
                    suite = this.suits[i];
                    if (this.in_suit(j, suite)) {
                        cnt++;
                    }
                }

                if (cnt == z) {
                    return j;
                }
            }
        }

        return 0;
    };

    /**
     * Ten, Jack, Queen, King, Ace, in same suit
     */
    Hand.prototype.scoreRoyalFlush = function () {
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
                return 1;
            }
        }

        return 0;
    };

    /**
     * All cards are consecutive values of same suit
     */
    Hand.prototype.scoreStraightFlush = function () {
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
                    return max;
                }
            }
        }

        return 0;
    };

    /**
     * Four cards of the same value
     */
    Hand.prototype.scoreFourOfKind = function () {
        return this.scoreSameValue(4);
    };

    /**
     * Three of a kind and a pair
     */
    Hand.prototype.scoreFullHouse = function () {
        var three = this.scoreSameValue(3);

        if (three && this.scoreSameValue(2)) {
            return three;
        }

        return 0;
    };

    /**
     * All cards of the same suit
     */
    Hand.prototype.scoreFlush = function () {
        var max, suite;
        for (var i = 0; i < 4; i++) {
            suite = this.suits[i];
            if (this[suite].length == 5) {
                max = this[suite].reduce(function (prev, curr) {
                    return prev < curr ? curr: prev;
                }, 0);

                return max;
            }
        }

        return 0;
    };

    /**
     * All cards are consecutive values
     */
    Hand.prototype.scoreStraight = function () {
        var arr = [];
        var min, max, suite;

        for (var i = 0; i < 4; i++) {
            suite = this.suits[i];
            arr = arr.concat(this[suite]);
        }

        function onlyUnique(value, index, self) {
            return self.indexOf(value) === index;
        }

        if (arr.filter(onlyUnique).length == 5) {
            max = arr.reduce(function (prev, curr) {
                return prev < curr ? curr: prev;
            }, 0);

            min = arr.reduce(function (prev, curr) {
                return prev > curr ? curr: prev;
            }, 15);

            if (max == min + 4) {
                return max;
            }
        }

        return 0;
    };

    /**
     * Three of a Kind
     */
    Hand.prototype.scoreThreeOfKind = function () {
        return this.scoreSameValue(3);
    };

    /**
     * Two Pairs
     */
    Hand.prototype.scoreTwoPairs = function () {
        var two = this.scoreSameValue(2);
        if (two && this.scoreSameValue(2, two)) {
            return two;
        }

        return 0;
    };

    /**
     * One Pair
     */
    Hand.prototype.scoreOnePair = function () {
        return this.scoreSameValue(2);
    };

    /**
     * High Card
     */
    Hand.prototype.scoreHighCard = function () {
        return this.scoreSameValue(1);
    };

    Hand.prototype.total = function () {
        var total = this.scoreHighCard();

        total += this.scoreOnePair() * 15;
        total += this.scoreTwoPairs() * 225;
        total += this.scoreThreeOfKind() * 3375;
        total += this.scoreStraight() * 50625;
        total += this.scoreFlush() * 759375;
        total += this.scoreFullHouse() * 11390625;
        total += this.scoreFourOfKind() * 170859375;
        total += this.scoreStraightFlush() * 2562890625;
        total += this.scoreRoyalFlush() * 38443359375;

        return total;
    };

    function checkHand(string) {
        var cards = string.split(' ');
        var hand1 = new Hand(cards[0], cards[1], cards[2], cards[3], cards[4]);
        var hand2 = new Hand(cards[5], cards[6], cards[7], cards[8], cards[9]);

        return (hand1.total() > hand2.total());
    }

    fs.readFile('./p054_poker.txt', {encoding: 'ascii'}, function (err, data) {
        if (err) {
            throw err;
        }

        var str, win = 0;
        for (var i = 0; i < datalength; i++) {
            str = '';
            for (var j = 0; j < 30; j++) {
                if (j < 29) {
                    str += data[30 * i + j];
                }
            }
            if (checkHand(str)) {
                win++;
            }
        }

        console.log(win);
    });
})();
