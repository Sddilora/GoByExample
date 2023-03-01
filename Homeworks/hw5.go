/*Bir dizideki en küçük elemanı ekrana yazdıran algoritmayı ve akış şemasını hazırlayınız*/

package main

import "fmt"

func main() {

	a := []int{2, 5, 7, 6, 10, 4, 8, 0, 9, -1, 11, 3}
	//var kucuk int
	kucuk := a[0]

	for i := 0; i < len(a); i++ {

		if a[i] <= kucuk {
			kucuk = a[i]
		} else {
			//kucuk = kucuk
		}

		//if kucuk == a[0] {
		//	fmt.Println(a[0])

	}
	fmt.Println(kucuk)

}

/*if a[i] < kucuk {
	kucuk = a[i]
} else if a[i] > kucuk {
	continue
} else {
	kucuk = a[i]
}
fmt.Println(kucuk)*/
