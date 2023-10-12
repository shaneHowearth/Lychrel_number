// Package main -
// https://en.wikipedia.org/wiki/Lychrel_number
// A Lychrel number is a natural number that cannot form a palindrome through
// the iterative process of repeatedly reversing its digits and adding the
// resulting numbers. This process is sometimes called the 196-algorithm, after
// the most famous number associated with the process. In base ten, no Lychrel
// numbers have been yet proven to exist, but many, including 196, are suspected
// on heuristic[1] and statistical grounds. The name "Lychrel" was coined by
// Wade Van Landingham as a rough anagram of "Cheryl", his girlfriend's first
// name.
package main

import (
	"fmt"
	"strconv"
)

func getReversed(input string) string {

	rns := []rune(input)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func addTwoStrings(input, reversed string) string {
	inR := []rune(input)
	revR := []rune(reversed)
	carry := false
	output := make([]rune, 0, len(inR)+1) // output can never be longer than the length of input, plus one for a carry
	extra := rune(0)
	for i, j := len(inR)-1, len(revR)-1; i >= 0; i, j = i-1, j-1 {
		// fmt.Printf("Adding: %s, (%d) + %s (%d)\n", string(inR[i]), inR[i], string(revR[j]), revR[j])
		// Is there anything left in revRune to use?
		revRune := rune(48)
		correction := rune(0)
		if j >= 0 {
			revRune = revR[j]
			correction = rune(48)
		}

		// The runes for numbers 0 - 9 are 48 - 57
		// So instead of an expensive string manipulation we just add the runes
		// together.

		// Carry one from the previous iteration.
		extra = rune(0)
		if carry {
			extra = rune(1)
		}

		// The actual addition.
		// Add the rune in input, the corresponding rune in reverse rune (if
		// there is one) and any extra carried over from the previous iteration.
		// Minus rune(48)/zero from the value to bring the output back to 0 - 9.
		// Note if revRune is > 48 then we need to correct the calculation
		x := inR[i] + revRune + extra - correction

		// Will one be carried through to the next iteration.
		carry = x > 57
		if carry {
			x -= rune(10)
		}
		if x > 57 {
			x %= rune(57)
		}

		output = append(output, x)
	}

	if carry {
		output = append(output, rune(49))
	}

	// fmt.Println("Output: ", string(output), output)
	return getReversed(string(output))
}

func lychrel(i int) {
	input := strconv.Itoa(i)
	for {
		fmt.Println("Input:", input)
		reversed := getReversed(input)
		if reversed == input {
			// w00t!!!
			fmt.Println("Found:", reversed)
			return
		}
		// Adding the ints in the two strings :(
		input = addTwoStrings(input, reversed)
	}
}

func main() {
	for i := 10; i < 196; i++ {
		fmt.Println("####################")
		fmt.Println("Finding: ", i)
		lychrel(i)
	}
}
