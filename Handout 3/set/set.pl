%%Split up functions into different cases

%%isUnion(Set1,Set2,Union) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.

isUnion([], Y, Y).
isUnion([H|T], Y, Z):-
    check_member(H, Y),!,isUnion(T,Y,Z).
isUnion([H|T],Y,[H|Z]):-
    \+check_member(H,Y),isUnion(T,Y,Z).
%%May be another case I cant think of. DOUBLE CHECK

%%isIntersection(Set1,Set2,Intersection) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
isIntersection([], _, []).
isIntersection([F|R1], Y, I) :-
  check_member(F, Y),!,I = [F|R3],
  isIntersection(R1, Y, R3).
isIntersection([_|R1], Y, I) :-
  isIntersection(R1, Y, I).

isEqual(X, Y) :- %%simple check for intersection and equal
    %% remove fail and add body/other cases for this predicate
    %%fail.
    isIntersection(X, Y, Z), Z=X.

%% Helper Function
check_member(I, [I|_]).
check_member(I, [_|R]) :-
    check_member(I, R).
