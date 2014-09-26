;; Distinct powers
;;
;; Completed on Fri, 26 Sep 2014, 23:25

(def maxa 100N)
(def maxb 100)

(defn fpow [x n]
  (if (= n 1)
    x
    (* x (fpow x (- n 1)))))

(println (count (distinct
                 (for [i (range 2N (inc maxa)) j (range 2 (inc maxb))]
                   (fpow i j)))))
