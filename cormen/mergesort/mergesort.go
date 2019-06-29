//mergesort based on cormen p29-32, 2nd edition
package mergesort

import (
	"math"
)

const (
	// maximum possible integer used as the sentinel
	sentinel = math.MaxInt64
)

//mergesort is an implementation of mergesort algorithm in corment p29.32
func mergesort(cards []int) {
	//merge combines two sorted sub arrays into a sorted array
	//start is the start of the left sub array
	//mid is beginning for the right sub array
	//end is end of the right sub array
	var merge = func(start, mid, end int) {
		n1, n2 := mid-start, end-mid //counts of how many entries in two sub arrays

		// create two separate piles of the two sorted sub arrays
		L, R := make([]int, n1, n1+1), make([]int, n2, n2+1)
		copy(L[0:n1], cards[start:mid])
		copy(R[0:n2], cards[mid:end])

		// add the sentinel to the end of each so you never need to test if a pile is empty
		L, R = append(L, sentinel), append(R, sentinel)
		defer func() { L, R = nil, nil }() //throw away sentinels when sorting is finished

		// because you know exactly how many cards there are the sentinel itself never gets added
		// pick the exact number of cards up you need always taking the smallest top value card
		// and putting that in the first place in the final sorted array
		// the sentinels never get picked up
		i, j := 0, 0
		for k := start; k < end; k++ { // this runs in linear time
			if L[i] <= R[j] {
				cards[k] = L[i]
				i++
				continue
			}
			cards[k] = R[j]
			j++
			continue
		}
	}

	var sort func(start, end int) // needs to be pre-declared as its a recursive function
	// divide and conquer recursive
	// because this divides into two each recursion it will be called lg(n) times.
	sort = func(start, end int) {
		//if the beginning meets the end then its obviously sorted and nothing to do
		//this is the base recursive case
		if start < end-1 {
			mid := int(math.Floor(float64((start + end) / 2))) //find the median card
			sort(start, mid)                                   //sort the left side
			sort(mid, end)                                     //sort the right side
			merge(start, mid, end)                             //merge the two sorted piles which takes linear time
		}
	}

	//sort the entire pack
	sort(0, len(cards))

	//as sort gets called recursively lg (n) times and each of those calls, calls merge which takes (n) time
	//mergesort runs in [n lg (n)] time or O(n lg n)
}

//mergesort2 is a rewrite of mergesort to not use sentinels as per exercise 2.2-3 Cormen, 2nd Edition, p36
//Rewrite the MERGE procedure so that it does not use sentinels,
//instead stopping once either array L or R has had all of its elements copied back to A and then copying
//the remainder of the other array back into A
func mergesort2(cards []int) {
	//merge combines two sorted sub arrays into a sorted array
	//start is the start of the left sub array
	//mid is beginning for the right sub array
	//end is end of the right sub array
	var merge = func(start, mid, end int) {
		nl, nr := mid-start, end-mid //counts of how many entries in two sub arrays

		// create two separate piles of the two sorted sub arrays
		L, R := make([]int, nl), make([]int, nr, nr+1)
		copy(L[0:nl], cards[start:mid])
		copy(R[0:nr], cards[mid:end])
		
		l, r := 0, 0
		for card := start; card < end; card++ { // this runs in linear time
			switch {
			case l-nl == 0: // L is empty
				copy(cards[card:end], R[r:])
				return
			case r-nr == 0: // R is empty
				copy(cards[card:end], L[l:])
				return
			default:
				if L[l] <= R[r] {
					cards[card] = L[l]
					l++
					continue
				}
				cards[card] = R[r]
				r++
				continue
			}
		}
	}

	var sort func(start, end int) // needs to be pre-declared as its a recursive function
	// divide and conquer recursive
	// because this divides into two each recursion it will be called lg(n) times.
	sort = func(start, end int) {
		//if the beginning meets the end then its obviously sorted and nothing to do
		//this is the base recursive case
		if start < end-1 {
			mid := int(math.Floor(float64((start + end) / 2))) //find the median card
			sort(start, mid)                                   //sort the left side
			sort(mid, end)                                     //sort the right side
			merge(start, mid, end)                             //merge the two sorted piles which takes linear time
		}
	}
	//sort the entire pack
	sort(0, len(cards))
	//as sort gets called recursively lg (n) times and each of those calls, calls merge which takes (n) time
	//mergesort runs in [n lg (n)] time or O(n lg n)
}
