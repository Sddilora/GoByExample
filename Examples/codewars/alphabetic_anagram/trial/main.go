package main

import (
	"fmt"
	"math/big"
	"sort"
)

func main() {

	a := "SEDA"
	fmt.Println(ListPosition(a))
}

func ListPosition(word string) *big.Int {

	var wordsOrder *big.Int
	repeatedLetters := make(map[string]int)

	stringArr := []string{}
	stringArrSorted := []string{}

	for i := 0; i < len(word); i++ {
		for c := 0; c < len(word); c++ {
			if i != c && word[i] == word[c] {
				repeatedLetters[string(word[i])] += 1
				fmt.Println(repeatedLetters, c, i)
				break
			}
		}
	}

	for i := 0; i < len(word); i++ {
		stringArr = append(stringArr, string(word[i]))
		stringArrSorted = append(stringArrSorted, string(word[i]))
	}

	sort.Strings(stringArrSorted) // now stringArrSorted variable updated as sorted alphabetically

	for i := 0; i < len(word); i++ {
		for c := 0; c < len(word); c++ {
			if len(repeatedLetters) == 0 {
				if stringArr[i] == stringArrSorted[c] {
					stringArrSorted = RemoveIndex(stringArrSorted, c)
					wordsOrder.Mul(factorial(int64(len(stringArrSorted))), (intToBigInt(c)))
				}
				fmt.Println(0)
			} else {
				fmt.Println("this part will be handled")
			}
		}
	}

	fmt.Println(repeatedLetters, wordsOrder)

	return wordsOrder
}

func factorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	bigN := big.NewInt(n)
	one := big.NewInt(1)
	result := new(big.Int)
	result.Mul(bigN, factorial(n-1))
	result.Mul(result, one)

	return result
}

func intToBigInt(n int) *big.Int {
	result := big.NewInt(0)
	result.SetInt64(int64(n))
	return result
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
