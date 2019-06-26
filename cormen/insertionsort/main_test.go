package insertionsort

import (
	"fmt"
	"testing"
)

func TestIsort(t *testing.T) {
	tests := []struct {input []int; want []int}{
		{[]int{5, 2, 4, 6, 1, 3}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _,test := range tests {
		isort(test.input)
		fmt.Println(test.input)
	}
}
