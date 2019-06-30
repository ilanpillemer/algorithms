package insertionsort

import (
	"fmt"
	"testing"
)

func TestIsort(t *testing.T) {
	tests := []struct {input []int}{
		{[]int{5, 2, 4, 6, 1, 3}},
		{[]int{5, 2, 4, 6, 1, 3}},
		{[]int{5, 2, 4, 6, 1, 3, 3, 7, 8, 2, 6, 5, 3, 1, 6, 8, 9, 10, 54, 1, 5, 8}},
		{[]int{5}},
		{[]int{}},
		{[]int{5, 4, 2, 1, 4}},
		{[]int{5, 4, 2, 1, 4, 7, 3, 5, 6, 345, 567, 234, 765, 3245, 6574, 4536, 6785, 54, 56, 76, 12, 53, 83, 12, 7, 3, 675467, 0, 23423, -5, 2341, -12324, 234, 4563, -4536}},
	}
	for _,test := range tests {
		isort(test.input)
		if !isSorted(test.input) {
			t.Errorf("%v",test.input)
		}
//		fmt.Println(test.input)
	}
}

func TestIsortR(t *testing.T) {
	tests := []struct {input []int}{
		{[]int{5, 2, 4, 6, 1, 3}},
		{[]int{5, 2, 4, 6, 1, 3}},
		{[]int{5, 2, 4, 6, 1, 3, 3, 7, 8, 2, 6, 5, 3, 1, 6, 8, 9, 10, 54, 1, 5, 8}},
		{[]int{5}},
		{[]int{}},
		{[]int{5, 4, 2, 1, 4}},
	}
	for _,test := range tests {
		isortR(test.input)
		if !isSorted(test.input) {
			t.Errorf("%v",test.input)
		}
		fmt.Println(test.input)
	}
}


func isSorted(input []int) bool {
	for i := range input {
		if i == 0 {
			continue
		}
		if input[i-1] > input [i] {
			return false
		}
	}
	return true
}