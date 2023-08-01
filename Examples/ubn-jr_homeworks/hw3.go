package hw

import "fmt"

func hw3() {

	toplam := 0

	for b := 0; b <= 100; b++ {
		toplam = b + toplam
	}
	fmt.Println(toplam)
}

/*:1’den 100’e kadar olan sayıların toplamını ekrana yazdıran algoritmayı ve akış şemasını
hazırlayınız.*/
