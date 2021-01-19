year_1953_1996_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    %%fail
    novel(Book,1953); novel(Book,1996).

period_1800_1900_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    novel(Book,Year), Year >= 1800, Year =< 1900.

lotr_fans(Fan) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    fan(Fan, Books), member(the_lord_of_the_rings, Books).

author_names(Author) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    fan(chandler, List1), author(Author, List2), list_compare(List1, List2).

fans_names(Fan) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    fan(Fan, List1), author(brandon_sanderson, List2), list_compare(List1, List2).

mutual_novels(Book) :-
    %% remove fail and add body/other cases for this predicate
    %%fail.
    common(phoebe, ross, Book);
    common(phoebe, monica, Book);
    common(ross, monica, Book).

%% Helper Function
    %% Function to compare elements of first list to second list
    list_compare([H|T],L) :- 
        member(H,L).
    list_compare([H|T],L) :- 
        not(member(H,L)), list_compare(T,L).
    
    %% Function to find commonalities between lists
common(First, Second, Book) :-
    fan(First, List1), fan(Second, List2), member(Book, List1), member(Book, List2).