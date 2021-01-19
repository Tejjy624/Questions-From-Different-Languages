(defun pivot (n xs)
  (list (smaller n xs) (larger n xs))
)

(defun quicksort (xs)
  (cond
    ((null xs) nil)
    (t
      (let* ((pivnum (car xs)) (currlist (pivot pivnum xs)))
        (cond ;;go through elements and reorder for every element
          (
            (and (null (car currlist)) (null (cadr currlist)))
            nil
          )
          (
            (null (car currlist)) (append (list pivnum) (quicksort (cdadr currlist)))
          )
          (
            (null (cadr currlist)) (append (quicksort (car currlist)) (list pivnum))
          )
          (t
            (append (quicksort (car currlist)) (list pivnum) (quicksort (cdadr currlist)))
          )
        )
      )
    )
  )
)

;;  Makes the list with numbers smaller than n
(defun smaller (n xs)
  (cond
    ((null xs) nil)
    ((< (car xs) n) (append (list (car xs)) (smaller n (cdr xs)) ))
    (t (smaller n (cdr xs)))
  )
)

;; Makes the list with numbers greater than n
(defun larger (n xs)
  (cond
    ((null xs) nil)
    ((>= (car xs) n) (append (list (car xs)) (larger n (cdr xs)) ))
    (t (larger n (cdr xs)))
  )
)