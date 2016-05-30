/**
 * Odd period square roots
 *
 * Completed on
 */

import java.lang.Math;

public class Euler064 {
    class Elem {
        int base;
        double epsilon, rest;

        Elem(double x, double eps) {
            base = (int)x;
            rest = x - (double)base;

            if (rest > 0.0d) {
                epsilon = eps / (rest * rest);
            }
        }
    }

    void iteration(int i) {
        double t = Math.sqrt(i * 1.0d);
        double eps = 1.0e-15;

        System.out.println("i: " + i);
        System.out.println("--------------------------------");

        stop:
        while (eps < 0.01d) {
            Elem el = new Elem(t, eps);
            System.out.println("b: " + el.base + "\te: " + el.epsilon);

            if (el.rest > 0.0d) {
                t = 1.0d / el.rest;
                eps = el.epsilon;
            } else {
                break stop;
            }
        }
        System.out.println("--------------------------------");
    }

    void solve() {
        int s;
        for (int i = 2; i < 10000; i++) {
            s = (int)Math.sqrt(i * 1.0d);
            if (i > s * s) {
                iteration(i);
            }
        }
    }

    public static void main(String[] args) {
        Euler064 eu = new Euler064();
        eu.solve();
    }
}
