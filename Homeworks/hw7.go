/*:Bir öğrencinin almış olduğu vize notu ile final notunun ortalamasını hesaplayan
algoritmayı ve akış şemasını hazırlayınız.(Vize notunun %30’unu final notunun %70’ini alarak
hesaplama yapınız) (Notlar kullanıcı tarafından girilecektir*/

package main

import "fmt"

func main() {

	vize := 0
	final := 0
	//ortalama := 0.0

	fmt.Println("Vize notunuzu giriniz")
	fmt.Scanf("%d\n", &vize)

	fmt.Println("Final notunuzu giriniz")
	fmt.Scanf("%d\n", &final)

	ortalama := (float64(vize) * 0.3) + (float64(final) * 0.7)

	fmt.Println("Ortalamanız", ortalama)

}
