//insertion sort, cormen p17
//Insertion Sort works the way you can sort cards in two hands.
//The right hand holds the unsorted Cards.
//The left Hand starts empty.
//You then take one card at a time from the Right Hand
//And you insert it into the correct position in the cards in the Left Hand
package insertionsort

//isort is an inplace implementaton of insertion sort derived from cormen p17 2nd edition
func isort(cards []int) {
	var insert = func(card int) {
		prev, value := card - 1, cards[card]
		//cards[0:prev] are the left hand, and hence always already sorted
		//go through the already sorted cards moving them one space to the right
		//until the left hand card is smaller or equal to the current card
		for ; prev > -1 && cards[prev] > value; prev-- { // sentinel is -1 as the first card is position 0.
			cards[prev+1] = cards[prev] // the previous card was higher so it should be shifted right leaving a space
		}
		cards[prev+1] = value // all cards to left are smaller so put card in here
	}

	// card is the "current card" being inserted into the hand
	// as the first card is already sorted as its only one card
	// one can assume first card is in the left hand already and
	// start process from card 1 (position 2)
	for card:=1; card < len(cards); card++ {
		// cards[:j] are always already sorted
		insert(card)
	}
}

