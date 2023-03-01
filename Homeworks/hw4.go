/*:1’den 100’e kadar olan tek sayıları ekrana yazdıran algoritmayı ve akış şemasını
hazırlayınız.*/

package main

import "fmt"

func main() {

	for a := 0; a < 100; a++ {

		if a%2 != 0 {
			fmt.Println(a)
		}
	}
}
