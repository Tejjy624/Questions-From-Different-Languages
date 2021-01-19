
(defun reachable (transition start final input)
	"Return true if the NFA accepts the input and returns false if it cannot be reachable"
    (cond
    	((listp start) (setq next (funcall transition (car start) (car input))))
    	(t (setq next (funcall transition start (car input)))))

    (cond
    	((and (eql start final)) t)
    	((or 
    		(reachable transition (car next) final (cdr input))
    		(reachable transition (cdr next) final (cdr input))) t)
    	(t nil)
    )
)