; Lattice paths

(defvar mem-cache (make-array '(21 21) :initial-element nil))

(defun z (i j)
  (if (aref mem-cache i j)
      (aref mem-cache i j)
      (let (x)
        (setq x
          (if (or (= 0 i) (= 0 j))
              1
              (+ (z (- i 1) j) (z i (- j 1)))))
	(setf (aref mem-cache i j) x))))

(print (z 20 20))
