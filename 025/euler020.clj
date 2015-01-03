;; Factorial digit sum
;;
;; Completed on Tue, 16 Sep 2014, 22:55

(def n 100N)

(defn fact [x]
  (if (= x 1)
    1
    (* x (fact (- x 1)))))

(defn to_list [x s]
  (if (zero? x)
    s
    (recur (bigint (/ x 10)) (conj s (mod x 10)))))

(println (reduce + (to_list (fact n) '())))
