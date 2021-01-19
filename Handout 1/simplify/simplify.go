package simplify

import "hw1/expr"

// Depth should return the simplified expresion
func Simplify(e expr.Expr, env expr.Env) expr.Expr {
	//TODO implement the simplify
	//panic("TODO: implement this!")
	switch e.(type) {
	//Dont need to simplify literal, its already simplified!
	case expr.Var: //Var in the environment, replace with its value
		n := e.(expr.Var)
		//check := env[n]
		if _, exist := env[n]; exist {
			//return expr.Literal(env[n])
			return expr.Literal(e.Eval(env))
		}
		return e
	case expr.Unary: //Unary with operand literal, replace with literal respresenting result
		n := e.(expr.Unary)
		n.X = Simplify(n.X, env)
		if _, check := n.X.(expr.Literal); check {
			return expr.Literal(n.Eval(env))
		}
		return n
	case expr.Binary: // 3 posibilities
		n := e.(expr.Binary)
		// 1st: Operands are literal, replace with literal
		if _, check := n.X.(expr.Literal); check {
			if _, check := n.Y.(expr.Literal); check {
				return expr.Literal(n.Eval(env))
			}
		}
		n.X = Simplify(n.X, env) 
		n.Y = Simplify(n.Y, env)
		
		// 2nd: One operand is 0, return only non-zero operand
		if _, check := n.X.(expr.Literal); !check {
			if _, check := n.Y.(expr.Literal); !check{
				return n
			} else if _, check := n.Y.(expr.Literal); check {
				if n.Y.Eval(env) == 0 {
					if n.Op == '+' {
						return n.X
					} else if n.Op == '*' {
						return expr.Literal(0)
					} else {
						return n
					}
				}else if n.Y.Eval(env) == 1 {
					if n.Op == '*' {
						return n.X
					} else {
						return n
					}
				} else {
					return n
				}
			} else {
				return n
			}
		} else if _,check := n.Y.(expr.Literal); !check {
			if _, check := n.X.(expr.Literal); !check{
				return n
			} else if _, check := n.X.(expr.Literal); check {
				if n.X.Eval(env) == 0 {
					if n.Op == '+' {
						return n.Y
					} else if n.Op == '*' {
						return expr.Literal(0)
					} else {
						return n
					}
				} else if n.X.Eval(env) == 1 {
					if n.Op == '*' {
						return n.Y
					} else {
						return n
					}
				} else {
					return n
				}
			} else {
				return n
			}
		}
		return expr.Literal(n.Eval(env))
	}
	return e
}
