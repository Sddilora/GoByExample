// Build a pyramid-shaped tower, as an array/list of strings, given a positive integer number of floors. A tower block is represented with "*" character.
package main

import "strings"

func main() {
	pyramid := TowerBuilder(5)
	for _, k := range pyramid {
		print(k, "\n")
	}
}

func TowerBuilder(nFloors int) []string {
	var pyramid []string
	for i := 0; i < nFloors; i++ {

		layer := strings.Repeat(" ", nFloors-1-i) + strings.Repeat("*", (2*i)+1) + strings.Repeat(" ", nFloors-1-i)
		pyramid = append(pyramid, layer)
	}

	return pyramid
}
