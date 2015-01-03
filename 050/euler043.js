/**
 * Sub-string divisibility
 *
 * Completed on Sat, 3 Jan 2015, 18:57
 */

(function () {
    function check_substring(arr) {
        var a1 = 100 * arr[1] + 10 * arr[2] + arr[3];
        if (a1 % 2 != 0) {
            return false;
        }

        var a2 = 100 * arr[2] + 10 * arr[3] + arr[4];
        if (a2 % 3 != 0) {
            return false;
        }

        var a3 = 100 * arr[3] + 10 * arr[4] + arr[5];
        if (a3 % 5 != 0) {
            return false;
        }

        var a4 = 100 * arr[4] + 10 * arr[5] + arr[6];
        if (a4 % 7 != 0) {
            return false;
        }

        var a5 = 100 * arr[5] + 10 * arr[6] + arr[7];
        if (a5 % 11 != 0) {
            return false;
        }

        var a6 = 100 * arr[6] + 10 * arr[7] + arr[8];
        if (a6 % 13 != 0) {
            return false;
        }

        var a7 = 100 * arr[7] + 10 * arr[8] + arr[9];
        if (a7 % 17 != 0) {
            return false;
        }

        var i, val = 0;
        var len = arr.length;

        for (i = 0; i < len; i++) {
            val *= 10;
            val += arr[i];
        }

        sum += val;
    }

    function permutation(inarray, outarray) {
        var i;
        var len = inarray.length;
        var tmp1, tmp2;

        if (len == 0) {
            check_substring(outarray);
        } else {
            for (i = 0; i < len; i++) {
                tmp1 = inarray.slice(0);
                tmp2 = outarray.slice(0);

                tmp2.push(tmp1[i]);
                tmp1.splice(i, 1);

                permutation(tmp1, tmp2);
            }
        }
    }

    var list = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
    var sum = 0;

    permutation(list, []);

    console.log(sum);
})();
