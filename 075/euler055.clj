;; Lychrel numbers
;;
;; Completed on

(require '[clojure.string :as s])

(defn palindrom?
  [x]
  (let [y (str x)]
    (= y (s/reverse y))))

(defn reverse_summ
  [x]
  (+ x (bigint (s/reverse (str x)))))
