;; Self powers
;;
;; Completed on Sat, 20 Sep 2014, 20:57

(def lim 1000)

(def last_num 10000000000)

(defn selfpow
  [x n]
  (if (= n 1)
    x
    (mod (* x (selfpow x (- n 1))) last_num)))

(defn dm
  [x]
  (selfpow x x))

(println (mod (reduce + (map dm (range 1 (+ lim 1)))) last_num))
