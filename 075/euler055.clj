;; Lychrel numbers
;;
;; Completed on Mon, 6 Apr 2015, 23:02

(require '[clojure.string :as s])

(def iteration_limit 50)

(defn palindrom?
  [x]
  (let [y (str x)]
    (= y (s/reverse y))))

(defn reverse_summ
  [x]
  (+ x (bigint (s/reverse (str x)))))

(defn lychrel?
  [t, i]
  (if (= i 0)
    true
    (let [z (reverse_summ t)]
      (if (palindrom? z)
        false
        (recur z (- i 1))))))

(println (count (filter #(lychrel? % iteration_limit) (range 1N 10000N))))
