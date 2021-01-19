package nfa
import (
	"sync"
)
// A nondeterministic Finite Automaton (NFA) consists of states,
// symbols in an alphabet, and a transition function.

// A state in the NFA is represented as an unsigned integer.
type state uint

// Given the current state and a symbol, the transition function
// of an NFA returns the set of next states the NFA can transition to
// on reading the given symbol.
// This set of next states could be empty.
type TransitionFunction func(st state, act rune) []state

// Reachable returns true if there exists a sequence of transitions
// from `transitions` such that if the NFA starts at the start state
// `start` it would reach the final state `final` after reading the
// entire sequence of symbols `input`; Reachable returns false otherwise.
func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO
	//panic("TODO: implement this!")
	var wait sync.WaitGroup
	Channel := make(chan bool, 100000) //buffered channel
	wait.Add(1) //add wait to run function without issues. Should handle data races as well
	CheckReachability(transitions, start, final, input, &wait, Channel)
	wait.Wait()
	close(Channel)
	/*go func(){
	    wg.Wait()
	    close(returns)
	}()*/
	for i := range Channel {
		if i {
			return true
		}
	}
	return false
}

//Helper Function (split into cases)
func CheckReachability(
	transitions TransitionFunction,
	start,
	final state,
	input []rune,
	wait *sync.WaitGroup,
	Channel chan bool) {
	//No data races should be present
	if len(input) == 0 {
		if start == final {
			Channel <- true
			//return true
		} /*else {
			return false
		}*/
	} else {
		PossibleStates := transitions(start, input[0])
		//for _, state := range transitions(start,input[0]) {
		for _, state := range PossibleStates {
			/*if CheckReachability(transitions, state, final, input[1:len(input)], wait, Channel) {
				return true
			}*/
			wait.Add(1)
			go CheckReachability(transitions, state, final, input[1:len(input)], wait, Channel)
		}
	}
	Channel <- false
	wait.Done()
}