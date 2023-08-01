package defer_statement

import "fmt"

func A2(sayi int32) string {

	defer B2()

	if sayi%2 == 0 {
		return "Çift sayıdır"
	}
	if sayi%2 != 0 {
		return "Tek sayidir"
	}

	// eğer buraya B2 yi çağırırsam veya defer B2 şeklinde çağırırsam "Bitti" yazdırmaz fonksiyonum. ilk returne ulaşınca fonksiyondan çıkar.
	return "Belli Değil"
}

func TestA2() {
	sonuc := A2(9)
	fmt.Println(sonuc) //Sonuc A2 fonksiyonu bittikten sonra yazıldığı için önce B2 sonra sonuç yazdırılır. A2'nin içinde defer ile çağırılan B2 fonksiyon bitene kadar ertlenir fonksiyondan çıkıldığında yazdırlır.
}

func B2() {
	fmt.Println("Bitti")
}
