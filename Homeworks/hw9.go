/*
	Kullanıcı tarafından oluşturulan bir dizideki elemanları büyükten küçüğe doğru

sıralayan algoritmayı ve akış şemasını hazırlayını
*/
package main

import "fmt"

var c int
var b []int
var a []int // = []int{4, 1, 5, 3, 9, 7, 11, 13, 14, 15, 8}
var m int
var n int
var h int

//var smallest int = a[0]

func main() {

	fmt.Println("Gireceğiniz Dizinin Eleman Sayısını Giriniz")
	fmt.Scanf("%d\n", &h)
	fmt.Println("Dizinizi arasına virgül koyarak ya da her sayıdan sonra entera basarak Giriniz")
	for i := 0; i < h; i++ {

		fmt.Scanf("%d\n", &n)
		a = append(a, n)
	}

	//for b := 0; b < len(a)-1; b++ {
	for k := 0; k < len(a)-1; k++ {

		for i := 0; i < len(a)-1; i++ {

			if a[i] < a[i+1] {
				continue
			} else {
				m = a[i]
				a[i] = a[i+1]
				a[i+1] = m

			}
		}
	}

	fmt.Println(a)

}
