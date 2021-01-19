%%Split into different cases
reachable(StartState, FinalState, []) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    StartState == FinalState. %%Base case

reachable(StartState, FinalState, [Begin|End]) :-
    transition(StartState, Begin, Next), %%May need to import transitions
    %%[transitions]. from piazza
    member(Curr, Next),
    reachable(Curr, FinalState, End).