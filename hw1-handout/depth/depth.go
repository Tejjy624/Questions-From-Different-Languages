package depth

import "hw1/expr"

// Depth returns the maximum number of AST nodes between the
// root of the tree and any leaf (literal or variable) in the tree.
func Depth(e expr.Expr) uint {
	// TODO: implement this function
	//return 1234
	//var bracket uint = 0
	var depth uint = 1
	input := expr.Format(e)
	length := len(input)
	for _, char := range input[0:length] {//input[1:len(input)-1] {
		if char == ')' { //|| char == '-' {
			depth++
		}
		//if char =='(' {	//check brackets after loop
		//	if bracket == '0' {
		//		depth--
		//		bracket++
		//	}
		//}
		//if char == ')' {  //might need this to check end depth
			//if bracket == '1' {
			//	depth++
			//	bracket--
			//}
		//}
	}
	//check for brackets at end and adjust depth
	if input[len(input)-1] == ')' && input[len(input)-2] == ')' {
		return (depth -1)
	}
	//if input[len(input)-2] == ')' {
		//depth
	//}
	return depth
}
