;; Power digit sum
;;
;; Completed on Mon, 15 Sep 2014, 23:29
;;
;; java -cp ~/path/to/clojure-1.8.0.jar clojure.main euler016.clj

(def power 1000)

(defn fpow [x n]
  (if (= n 1)
    x
    (* x (fpow x (- n 1)))))

(defn to_list [x s]
  (if (zero? x)
    s
    (recur (bigint (/ x 10)) (conj s (mod x 10)))))

(println (reduce + (to_list (fpow 2N power) '())))
