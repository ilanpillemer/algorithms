package subsets

import (
	"fmt"
	"testing"
)

// shows all the subsets
func TestShow(t *testing.T) {
	f := func(subset []int) {
		fmt.Println(subset)
	}
	subsets(4, f)
}

// gets the sums of all the subsets
func TestSum(t *testing.T) {
	f := func(subset []int) {
		sum := 0
		for _,v := range subset {
			sum += v
		}
		fmt.Printf("sum of %v is %d\n",subset,sum)
	}
	subsets(4, f)
}

