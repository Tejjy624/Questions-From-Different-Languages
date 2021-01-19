;; Function to find the min
(defun minfunc (num list)
  (cond 
    ((null (car list)) num)  ;;list empty/ended
    ((< (car list) num) (minfunc (car list) (cdr list)))   ;;new min found
    (t (minfunc num  (cdr list))))   ;;min not changed
)

;; Function to find the mean
(defun meanfunc (xs total size)
    (cond
        ((null xs)
            (/ total size)
        )
        ((not (null xs))
            (setq newtotal (+ total (car xs)))
            (setq newsize (+ size 1))
            (setq newxs (cdr xs))
            (meanfunc newxs newtotal newsize)
        )
    )
)

;; Function to find the max
(defun maxfunc (xs max)
    (cond 
        ((null xs)  ;;list empty/ended
            (+ max 0)
        )
        ((not (null xs))
            (cond 
                ((< max (car xs))   ;;value higher than curr max
                    (maxfunc (cdr xs) (car xs))
                )
                ((or (> max (car xs)) (= max (car xs))) ;;value is lower/same as curr max
                    (maxfunc (cdr xs) max)
                )
            )
        )
    )
)

(defun min-mean-max (xs)
    (setq min (minfunc (car xs) (cdr xs)))
    (setq mean (meanfunc xs 0 0))
    (setq max (maxfunc (cdr xs) (car xs)))
    (list min mean max)
)