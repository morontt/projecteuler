/**
 * Largest palindrome product
 *
 * Completed on Tue, 9 Sep 2014, 00:08
 */

var i, j, k,
    max = 0;

function is_palindrome(t) {
    var len = t.length;

    if (len == 1 || len == 0) {
        return true;
    }

    return (t[0] == t[len - 1]) && is_palindrome(t.substr(1, len - 2));
}

for (i = 100; i < 1000; i++) {
    for (j = i; j < 1000; j++) {
        k = i * j;
        if (k > max && is_palindrome(k.toString())) {
            max = k;
        }
    }
}

console.log(max);
