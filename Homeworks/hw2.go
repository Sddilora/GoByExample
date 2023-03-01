package main

import "fmt"

func main() {

	//a:=0

	for a := 0; a < 100; a++ {

		if a%2 == 0 {
			fmt.Println(a, "bir çift sayıdır")
		} else {
			fmt.Println(a, "bir tek sayıdır")
		}
	}

}

/*:Bir sayının tek mi çift mi olduğunu ekrana yazdıran algoritmayı ve akış şemasını
hazırlayınız.*/
