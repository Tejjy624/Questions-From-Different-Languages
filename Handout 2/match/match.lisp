(defun match(pattern assertion)
   (let 
      ((pattop (car pattern))
         (asserttop (car assertion))
         (patremain (cdr pattern))
         (assertremain (cdr assertion))
      )
      (cond 
         ((and (null pattern) (null assertion)) ;; Both null
            t
         )
         ((or (null pattern) (null assertion)) ;; Only one is null
            nil
         )
         ((equal pattop asserttop)  ;; Both elements equal, go through all elements
            (match patremain assertremain)
         )
         ((equal pattop '?)   ;; Found "?", go through all elements
            (match patremain assertremain)
         )
         ((equal pattop '!)   ;; Found "!", 2 cases: followed by 1 or more pat/asserts
            (or (match patremain assertremain) (match pattern assertremain)
            )
         )
      )
   )
)