// benford package has functions to work with the Benford's Law
package benford

import (
	"fmt"
	"math"
)

// LawNumbers returns the probabilities of occurrence of each number from 1 to 9, based on the Benford's Law.
func LawNumbers() []Digit {
	var p []Digit
	l := len(BN) + 1
	for i := 1; i < l; i++ {
		log := math.Log10(1+(1/float64(i))) * 100
		p = append(p, Digit{Num: i, Prob: log})
	}

	return p
}

// Probabilities returns a the probabilities of occurrence of each number from 1 to 9, from a given set of values.
func Probabilities(data []float64) []Digit {
	ds := map[int]int{}
	t := 0
	for _, v := range data {
		ld := leadingDigit(int(v))
		if ld > 0 {
			ds[ld]++
			t++
		}
	}

	var p []Digit
	for k, v := range ds {
		per := (float64(v) * 100) / float64(t)
		p = append(p, Digit{Num: k, Prob: per})
	}

	return p
}

// ValidateLaw validates if a distribution fits the "Benford's law compliance theorem"
func ValidateLaw(observed, expected []Digit) {
	obs := ByNum(observed).ConvertToFloat()
	exp := ByNum(expected).ConvertToFloat()
	f := ChiSquaredTest(obs, exp)

	for i := 0; i < len(BN); i++ {
		fmt.Printf("%d:\tObserverd: %f\tExpected: %f\n", i+1, obs[i], exp[i])
	}

	fmt.Println("Chi-Square Value:", f)
}

// ChiSquaredTest returns the result of executing the Chi-Squared Test. It receives two parameters the observed values
// and the expected values.
func ChiSquaredTest(obs, exp []float64) float64 {
	var result float64
	for i, a := range obs {
		b := exp[i]
		if a == 0 && b == 0 {
			continue
		}
		result += (a - b) * (a - b) / b
	}
	return result
}
