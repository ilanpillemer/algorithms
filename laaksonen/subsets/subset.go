package subsets

//this is based on laaksonen p16, 1st edition
func subsets(n int, f func([]int)) {
	var subset = make([]int, 0)
	// this is a streaming algorithm for processing all the
	// subsets of set of n elements. The function expects
	// an array subset to exist in its closure
	var search func(int)
	search = func(k int) {
		if k == n+1 {
			// at a leaf so process
			f(subset)
			return
		}
		// include k in the subset
		subset = append(subset, k)
		search(k + 1)
		// dont include k in the subset
		subset = subset[:len(subset)-1]
		search(k + 1)
	}

	search(1)
}
