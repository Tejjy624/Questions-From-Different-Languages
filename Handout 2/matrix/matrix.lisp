(defun matrix-add (m1 m2)
    (mapcar #'(lambda (x y) (mapcar '+ x y)) m1 m2)
)

;; https://ml-cheatsheet.readthedocs.io/en/latest/linear_algebra.html
;; Used this website as reference for matrix manipulation
(defun matrix-multiply (a b)
    (defun rowmultiplication (r1 m2)    ;;choose individual rows and solve
        (cond
            ((null m2) nil) ;;end of matrix
            (t (cons (dotproduct r1 (car m2)) (rowmultiplication r1 (cdr m2))))
        )
    )
    (defun dotproduct (r1 r2)   ;;simple dot product calculation from the website
        (cond
            ((or (null r1) (null r2)) 0) ;;end of matrix
            (t (+ (* (car r1) (car r2)) (dotproduct (cdr r1) (cdr r2))))
        )
    )
    (defun multiply (m1 m2) ;;final multiplication
        (cond
            ((null m1) nil) ;;end of matrix
            (t (cons (rowmultiplication (car m1) m2) (multiply (cdr m1) m2)))
        )
    )
    (multiply a (matrix-transpose b))
)

(defun matrix-transpose (m)
    (defun selecttranspose (m cars cdrs) ;;take 
        (cond
            ((null m) (cons cars cdrs))
            (t (selecttranspose (cdr m) (append cars (list (caar m))) (append cdrs (list (cdar m)))))
        )
    )
    (cond   ;;check matrix endings before transpose
        ((null (car m)) nil)
        (t (cons (car (selecttranspose m '() '())) (matrix-transpose (cdr (selecttranspose m '() '())))))
    )
)