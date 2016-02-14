/**
 * Cyclical figurate numbers
 */
public class Euler061 {
    public static void main(String[] args) {
        int i = 0, r;
        do {
            r = i * (i + 1) / 2;
            if (r > 999 && r < 10000) {
                // System.out.println("Triangle: " + r + " i = " + i);
            }
            i++;
        } while (r < 10000);

        i = 0;
        do {
            r = i * i;
            if (r > 999 && r < 10000) {
                // System.out.println("Square: " + r + " i = " + i);
            }
            i++;
        } while (r < 10000);

        i = 0;
        do {
            r = i * (3 * i - 1) / 2;
            if (r > 999 && r < 10000) {
                // System.out.println("Pentagonal: " + r + " i = " + i);
            }
            i++;
        } while (r < 10000);

        i = 0;
        do {
            r = i * (2 * i - 1);
            if (r > 999 && r < 10000) {
                // System.out.println("Hexagonal: " + r + " i = " + i);
            }
            i++;
        } while (r < 10000);

        i = 0;
        do {
            r = i * (5 * i - 3) / 2;
            if (r > 999 && r < 10000) {
                // System.out.println("Heptagonal: " + r + " i = " + i);
            }
            i++;
        } while (r < 10000);

        i = 0;
        do {
            r = i * (3 * i - 2);
            if (r > 999 && r < 10000) {
                // System.out.println("Octagonal: " + r + " i = " + i);
            }
            i++;
        } while (r < 10000);
    }
}
