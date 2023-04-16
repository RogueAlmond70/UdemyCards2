package main

/*
print()
- Log out the contents of a deck of cards.

shuffle()
- Shuffles all the cards in a deck.

deal()
- Create a 'hand' of cards.

saveToFile()
- Save a list of cards to a file on the local machine.

newDeckFromFile()
- Load a list of cards from the local machine.



*/

// For this we will implement the Fisher Yates shuffle.

/*
Fisher-Yates Shuffle Algorithm:
Input: An array or list of elements

1. For i from n - 1 to 1, where n is the number of elements in the array:
 1. Generate a random integer j such that 0 <= j <= i
 2. Swap the element at index i with the element at index j

Output: The shuffled array or list of elements
*/
