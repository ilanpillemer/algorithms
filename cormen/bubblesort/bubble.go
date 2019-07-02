package bubblesort

//bubblesort as per cormen ex 2-2, 2nd edition, p38
//bubblesort is a popular sorting algorithm. It works by repeatedly swapping adjacent
//elements that are out of order
func bubblesort(cards []int) {
	for i := range cards { // runs n times and in each of those times it..
		for j := len(cards)-1; j>i; j-- { //runs 1 + 2 + 3 + 4 + 5 + 6 + .. n times approx n(n-1) /2 times.
			if cards[j] < cards[j-1] {
				cards[j],cards[j-1] = cards[j-1],cards[j]
			}
		}
	}
	// total worst case running time therefore is n ( n-1 ) / 2 times
	// so worst case is the same as insertion sort
	// however best case is also n ( n -1 ) / 2 for bubblesort whilst insertion sort takes less time on partially sorted cards.
}

//4 , 5, 6, 1, 2, 3
// i == 0 (first card)
// j == goes downwards from last card to i
// 3 < 2 ? no
// 2 < 1 ? no
// 1 < 6 ? yes -> 4, 5 ,1 ,6 ,2 ,3
// 1 < 5 ? yes -> 4, 1 ,5 ,6 ,2 ,3
// 1 < 4 ? yes -> 1, 4 ,5 ,6 ,2 ,3

// after first iteration lowest card has to be lowest in position
// it bubbled top "top"

// i == 1 (second card)
// j counts down to second card
// 3 < 2? no
// 2 < 6? yes -> 1, 4, 5, 2, 6, 3
// 2 < 5? yes -> 1, 4, 2, 5, 6, 3
// 2 < 4? yes -> 1, 2, 4, 5, 6, 3
// after second iteration lowest card is lowest from second card forward
// it bubbled

// i == 3 (third card)
// j counts down to third card
// 3 < 6? yes -> 1, 2, 4, 5, 3, 6
// 3 < 5? yes -> 1, 2, 4, 3, 5, 5
// 3 < 4? yes -> 1, 2, 3, 4, 5, 6
// after third iteraion lowest card is lowest from third forward

// i == 4 (fourth card)
// j counts down to fourth card
// 6 < 5? no
// 6 < 4? no

// i == 5 (fifth card)
// j counts down to fifth card
// 6 < 5? no

// i==6 and j==6 so algorithm end

// at each iteration the lowest card bubbles to the top
// so by the end it has to be sorted



