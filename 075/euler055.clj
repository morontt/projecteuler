;; Lychrel numbers
;;
;; Completed on Mon, 6 Apr 2015, 23:02
;;
;; java -cp ~/path/to/clojure-1.8.0.jar clojure.main euler055.clj

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
