package benford

import "sort"

// BN is an array with the Benford's numbers
var BN = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Digit struct contains one of the Benford's numbers and the occurrence probability of it.
type Digit struct {
	Num  int
	Prob float64
}

// ByNum is used to sort the Digit struct by Num.
type ByNum []Digit

func (n ByNum) Len() int           { return len(n) }
func (n ByNum) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByNum) Less(i, j int) bool { return n[i].Num < n[j].Num }

// CompleteSlice adds the missing Benford's numbers to the slice with a probability of 0. It's useful to run the
// Chi-Square test with two slices with the same length.
func (n ByNum) CompleteSlice() []Digit {
	if len(n) == len(BN) {
		return n
	}

	t := n // To avoid 'Assignment to method receiver doesn't propagate to other calls' warning.
	m := make(map[int]bool)

	for _, v := range t {
		m[v.Num] = true
	}

	var diff []int
	for _, i := range BN {
		if _, ok := m[i]; !ok {
			diff = append(diff, i)
		}
	}

	for _, d := range diff {
		t = append(t, Digit{Num: d, Prob: 0})
	}
	sort.Sort(t)
	return t
}

func (n ByNum) ConvertToFloat() []float64 {
	var c []float64
	for _, v := range n {
		c = append(c, v.Prob)
	}
	return c
}

// leadingDigit return the leading digit of a number.
func leadingDigit(v int) int {
	l := len(BN)
	for v > l {
		v = v / 10
	}
	return v
}
