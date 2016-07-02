;; Digit fifth powers
;;
;; Completed on Thu, 25 Sep 2014, 23:00
;;
;; java -cp ~/path/to/clojure-1.8.0.jar clojure.main euler030.clj

(defn fifthpow [x]
  (int (Math/pow x 5)))

(defn to_list [x s]
  (if (zero? x)
    s
    (recur (int (/ x 10)) (conj s (mod x 10)))))

(defn fifthsum? [x]
  (= x (reduce + (map fifthpow (to_list x '())))))

(println (reduce + (filter fifthsum? (range 2 354295)))) ;; 6 * 9^5 - max value
