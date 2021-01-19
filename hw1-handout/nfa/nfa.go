package nfa

// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO Write the Reachable function,
	// return true if the nfa accepts the input and can reach the final state with that input,
	// return false otherwise
	//panic("TODO: implement this!")

	//Check input length. Check possible states with a path (Using Reachable())
	//Piazza: If start not equal to final, and intput=nil, output should be false
	//length := len(input)
	//if (length ==0) {
	if len(input) == 0 {
		if start == final { //Piazza: output should be true
			return true
		}
	}
	//Active nfa's
	//if length > 0 {
	if len(input) > 0 {
		statechange := transitions(start, input[0])
		for _, state := range statechange {
			//if Reachable(transitions, state, final, input[1:length])
			if Reachable(transitions, state, final, input[1:len(input)]) {
				return true
			}
		}
	}
	return false //Consider tests not covered, return false always
}
