/**
 * Concealed Square
 *
 * Completed on Sun, 21 Feb 2016, 02:14
 */

public class Euler206 {
    long[] b = new long[250];

    boolean checkLo(Long q) {
        long w = q % 1000;

        return w > 799 && w < 900;
    }

    boolean checkHi(Long q) {
        String str = q.toString();
        int len = str.length();

        return str.charAt(0) == '1'
            && str.charAt(2) == '2'
            && str.charAt(4) == '3'
            && str.charAt(6) == '4'
            && str.charAt(8) == '5'
            && str.charAt(10) == '6'
            && str.charAt(12) == '7';
    }

    void solve() {
        long s, q;
        int idx = 0;

        for (int j = 1; j < 1000; j++) {
            s = j * 10 + 3;
            q = s * s;
            if (checkLo(q)) {
                b[idx] = s;
                idx++;
            }
            s = j * 10 + 7;
            q = s * s;
            if (checkLo(q)) {
                b[idx] = s;
                idx++;
            }
        }

        stop:
        for (int j = 0; j < idx; j++) {
            // 10000 * 10101^2 = 1020302010000 (pattern 1.2.3.4.5.7)
            // 10000 * 13893^2 = 1930154490000
            for (long k = 10101; k < 13893; k++) {
                q = 10000 * k * k + 2 * k * b[j] + b[j] * b[j] / 10000;
                if (checkHi(q)) {
                    System.out.println("n: " + 10 * (10000 * k + b[j]));
                    break stop;
                }
            }
        }
    }

    public static void main(String[] args) {
        Euler206 eu = new Euler206();
        eu.solve();
    }
}
