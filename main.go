// Command word-break-problem solves the word break problem.
// Given a sentence without space and dictionary, find a breakup
// of the sentence into the words if possible.
package main

import (
	"fmt"
	"strings"
)

var d map[string]bool // dictionary

func main() {
	d = map[string]bool{
		"dog":      true,
		"tiger":    true,
		"monkey":   true,
		"elephant": true,
	}

	// TODO: consider a dictionary that allows more than one way to make up
	// a sentence. Then return all the possible answers, and perhaps rank them.

	cache = make(map[string]bool)

	s := "monkeytigerdog"
	if hasWords(s) {
		fmt.Println(strings.Join(words, " "))
	} else {
		fmt.Println("Not separatable into words.")
	}
	fmt.Printf("cache %#+v \n", cache)

}

var r int // recursion depth

var cache map[string]bool // for dynamic programming or memoize
var words []string        // the sequence of words that make of a sentence

// hasWords returns true if there is a sequence of words that make up the sentence
func hasWords(sentence string) bool {
	// T(1) = O(1) work for base cases, cache hit or len = 0
	answer, inCache := cache[sentence]
	if inCache {
		return answer
	}
	r++
	if len(sentence) == 0 {
		r--
		return true
	}

	// n = len(sentence)
	// T(n) = k + T(n-1) + T(n-2) + T(n-3) + ... T(1) = O(n^2)

	// If we consider the hash for dictionary lookup is O(n)
	// T(n) = k + 1 + T(n-1) + 2 + T(n-2) + 3 + T(n-3) + ... (n-1) + T(1) = O(2*n^2) = O(n^2)
	for i := 1; i <= len(sentence); i++ {

		p := sentence[:i] // prefix
		s := sentence[i:] // suffix
		fmt.Println(strings.Repeat("-", r), p, "+", s)

		// At worst, the prefix is len(s) = len(sentence) - 1, so O(n-1) work
		// Time to check if prefix is in the dictionary, compute the hash of sequence of chars
		// can use a rolling hash to only do constant work for each appended character.
		if d[p] && hasWords(s) {
			fmt.Print(strings.Repeat("-", r), " found it\n\n")
			words = append(words, p)
			fmt.Println("words: ", strings.Join(words, " "))
			r--
			cache[sentence] = true
			return true
		}
	}

	fmt.Print(strings.Repeat("-", r), " did not find it\n\n")
	if len(words) > 0 {
		words = words[:len(words)-1]
	}
	fmt.Println("words: ", strings.Join(words, " "))
	cache[sentence] = false
	return false
}
