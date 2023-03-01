/*
Kullanıcı tarafından girilen kilo ve boy bilgilerine göre vücut kitle indeksini hesaplayan
aynı zamanda bu sonuca göre kişinin ideal kilo durumunu ekrana yazdıran algoritmayı ve akış
şemasını hazırlayınız.
Not-1: Vücut kitle indeksi = Kilo / Boy * Boy
Not-2: Kilo => kg, Boy => m
Not-3: Vücut kitle indeksi < 18,5 ise Zayıf
 18,5 < Vücut kitle indeksi < 24,9 ise Normal
 25 < Vücut kitle indeksi < 29,9 ise Fazla kilolu
 30 < Vücut kitle indeksi < 34,9 ise I. derece obez
 35 < Vücut kitle indeksi < 39,9 ise II. derece obez
 Vücut kitle indeksi > 40 ise III. derece obez
*/
package main

import "fmt"

func main() {

	kilo := 0.0
	boy := 0.0
	fmt.Println("Kilonuzu Giriniz")
	fmt.Scanf("%f\n", &kilo)

	fmt.Println("Boyunuzu Giriniz")
	fmt.Scanf("%f\n", &boy)

	bmi := kilo / (boy * boy)

	if bmi < 18.5 {
		fmt.Println("Dünya sağlık örgütü standartlarına göre kilonuz normalin altında, zayıfsınız")

	} else if bmi < 24.9 && bmi > 18.5 {
		fmt.Println("Dünya sağlık örgütü standartlarına göre kilonuz normal")

	} else if bmi < 29.9 && bmi > 25 {
		fmt.Println("Dünya sağlık örgütü standartlarına göre kilonuz normalin üstünde, fazla kilolusunuz")

	} else if bmi < 34.9 && bmi > 30 {
		fmt.Println("Dünya sağlık örgütü standartlarına göre kilonuz normalin üstünde, 1. derece obezsiniz")

	} else if bmi < 39.9 && bmi > 35 {
		fmt.Println("Dünya sağlık örgütü standartlarına göre kilonuz normalin üstünde, 2. derece obezsiniz")
	} else if bmi > 40.0 {
		fmt.Println("Dünya sağlık örgütü standartlarına göre kilonuz normalin üstünde, 3. derece obezsiniz")
	} else {
		fmt.Println("Girdiğiniz değerler doğru değil")
	}
}
