package hw

import "fmt"

func hw6() {

	a := []int{1, 6, 8, 5, 9, 12, 4, 5, 7, 8}

	toplam := 0

	for i := 0; i < len(a); i++ {

		toplam = toplam + a[i]
	}

	ortalama := float32(toplam) / float32(len(a))

	fmt.Println(ortalama)
}

/*Bir dizinin ortalamasını bulan algoritmayı ve akış şemasını hazırlayınız.*/
