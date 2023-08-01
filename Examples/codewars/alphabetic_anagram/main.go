// Consider a "word" as any sequence of capital letters A-Z (not limited to just "dictionary words"). For any word with at least two different letters, there are other words composed of the same letters but in a different order (for instance, STATIONARILY/ANTIROYALIST, which happen to both be dictionary words; for our purposes "AAIILNORSTTY" is also a "word" composed of the same letters as these two).

// We can then assign a number to every word, based on where it falls in an alphabetically sorted list of all words made up of the same group of letters. One way to do this would be to generate the entire list of words and find the desired one, but this would be slow if the word is long.

// Given a word, return its number. Your function should be able to accept any word 25 letters or less in length (possibly with some letters repeated), and take no more than 500 milliseconds to run. To compare, when the solution code runs the 27 test cases in JS, it takes 101ms.

// For very large words, you'll run into number precision issues in JS (if the word's position is greater than 2^53). For the JS tests with large positions, there's some leeway (.000000001%). If you feel like you're getting it right for the smaller ranks, and only failing by rounding on the larger, submit a couple more times and see if it takes.

// Python, Java and Haskell have arbitrary integer precision, so you must be precise in those languages (unless someone corrects me).

// C# is using a long, which may not have the best precision, but the tests are locked so we can't change it.

// Sample words, with their rank:
// ABAB = 2
// AAAB = 1
// BAAA = 4
// QUESTION = 24572
// BOOKKEEPER = 10743

package main

import (
	"fmt"
	"math/big"
	"sort"
)

func main() {
	a := "SEDA"
	ListPosition(a)
	fmt.Println(a[0])
}

func ListPosition(word string) *big.Int {

	fmt.Println("len", len(word), word)

	stringArr := []string{}
	stringArrSorted := []string{}
	wordsOrder := 0

	for i := 0; i < len(word); i++ {
		stringArr = append(stringArr, string(word[i]))
		stringArrSorted = append(stringArrSorted, string(word[i]))
	}

	sort.Strings(stringArrSorted) // now stringArrSorted variable updated as sorted alphabetically

	for i := 0; i < len(stringArr); i++ {
		for c := 0; c < len(stringArrSorted); c++ {
			if stringArr[i] == stringArrSorted[c] {
				fmt.Println(i, stringArr[i], c, stringArrSorted[c])
				stringArrSorted = RemoveIndex(stringArrSorted, c)
				wordsOrder += factorial(len(stringArrSorted)) * (c)
				fmt.Println("wordsOrder:", wordsOrder, "/factorial(len(stringArr)-1-i):", factorial(len(stringArr)-1-i), "/RemoveIndex():", stringArrSorted, "/stringArr", stringArr)
			}
		}
	}

	wordsOrder += 1

	fmt.Println("wordsOrder", wordsOrder)
	return nil
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
