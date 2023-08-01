// Given an n x n array, return the array elements arranged from outermost elements to the middle element, traveling clockwise.
// array = [[1,2,3],
//
//	[4,5,6],
//	[7,8,9]]
//
// snail(array) #=> [1,2,3,6,9,8,7,4,5]
package main

import "fmt"

func main() {
	snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(Snail(snailMap))
}

func Snail(snaipMap [][]int) []int {
	var snail []int
	if len(snaipMap) == 0 {
		return []int{}
	}

	for i := 0; i < len(snaipMap); i++ {
		if len(snaipMap[i]) == 0 {
			return []int{}
		} else if len(snaipMap[i]) != len(snaipMap) {
			return []int{}

		}

	}

	if len(snaipMap)%2 == 0 {
		for i := 0; i < len(snaipMap)/2; i++ {
			for f := i; f < len(snaipMap)-i; f++ {
				snail = append(snail, snaipMap[i][f])
			}
			for s := i + 1; s < len(snaipMap)-1-i; s++ {
				snail = append(snail, snaipMap[s][len(snaipMap)-1-i])
			}
			for t := i; t < len(snaipMap)-1-i; t++ {
				snail = append(snail, snaipMap[len(snaipMap)-1-i][len(snaipMap)-1-t])
			}
			for fo := i + 1; fo < len(snaipMap)-i; fo++ {
				snail = append(snail, snaipMap[len(snaipMap)-fo][i])
			}

		}
	} else if len(snaipMap)%2 == 1 {
		for i := 0; i < (len(snaipMap)+1)/2; i++ {
			for f := i; f < len(snaipMap)-i; f++ {
				snail = append(snail, snaipMap[i][f])
			}
			for s := i + 1; s < len(snaipMap)-1-i; s++ {
				snail = append(snail, snaipMap[s][len(snaipMap)-1-i])
			}
			for t := i; t < len(snaipMap)-1-i; t++ {
				snail = append(snail, snaipMap[len(snaipMap)-1-i][len(snaipMap)-1-t])
			}
			for fo := i + 1; fo < len(snaipMap)-i; fo++ {
				snail = append(snail, snaipMap[len(snaipMap)-fo][i])
			}
		}
	}

	return snail
}
