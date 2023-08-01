package interfaces

import "fmt"

func assertion(i interface{}) {
	sayi := i.(int) //Tip doğrulama(int doğrulması)
	// eğer sayi,ok dersek yakalayabildi mi yakalayamadı mı diye de bakarız
	fmt.Println(sayi)
}

func assertionWithOk(i interface{}) {
	sayi1, ok := i.(int) //tip dönüümü yap yapabiliyorsan okeye true at sayiyi ata else false zero

	fmt.Println(sayi1, ok)
}

func Demo4() {
	var something interface{} = 5 //bir sting gelirse hata verir.
	var somethingelse interface{} = "Dilo"
	assertion(something)
	assertionWithOk(somethingelse) //panic: interface conversion: interface {} is string, not int Hatası alırız
}
